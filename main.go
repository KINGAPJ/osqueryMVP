package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"osqueryMVP/api"
	"osqueryMVP/data"
	"osqueryMVP/osquery"
)

func main() {
	var db *sql.DB

	clientID, err := osquery.GetClientID()
	if err != nil {
		log.Fatal(err)
	}

	dsn := "osqueryuser:osquerypass@tcp(127.0.0.1:3307)/osquerydb"
	db, err = sql.Open("mysql", dsn+"?parseTime=true")

	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer db.Close()

	programs, err := osquery.FetchInstalledPrograms(clientID)
	if err != nil {
		log.Fatal(err)
	}
	if err := data.UpsertInstalledPrograms(db, programs); err != nil {
		log.Fatal(err)
	}

	info, err := osquery.FetchSystemInfo(clientID)
	if err != nil {
		log.Fatal(err)
	}
	if err := data.UpsertSystemInfo(db, info); err != nil {
		log.Fatal(err)
	}

	log.Println("Data upserted successfully.")

	handler := &api.APIHandler{
		DB:       db,
		ClientID: clientID,
	}
	http.HandleFunc("/latest_data", handler.LatestDataHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// curl -v http://localhost:8080/latest_data -o output.txt
