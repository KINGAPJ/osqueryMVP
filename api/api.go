package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"osqueryMVP/models"
)

type APIHandler struct {
	DB       *sql.DB
	ClientID string
}

func GetLatestSystemInfo(db *sql.DB, clientID string) (models.SystemInfo, error) {
	var info models.SystemInfo

	query := `SELECT client_id, hostname, cpu_brand, cpu_logical_cores, cpu_physical_cores, hardware_model, hardware_vendor, physical_memory, timestamp FROM system_info WHERE client_id = ? ORDER BY timestamp DESC LIMIT 1`
	err := db.QueryRow(query, clientID).Scan(
		&info.ClientID,
		&info.Hostname,
		&info.CPUBrand,
		&info.CPULogicalCores,
		&info.CPUPhysicalCores,
		&info.HardwareModel,
		&info.HardwareVendor,
		&info.PhysicalMemory,
		&info.Timestamp,
	)
	return info, err
}

func GetLatestInstalledApps(db *sql.DB, clientID string) ([]models.InstalledProgram, error) {
	query := `SELECT client_id, name, version, install_date, timestamp FROM installed_programs WHERE client_id = ? ORDER BY timestamp DESC LIMIT 1`

	rows, err := db.Query(query, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []models.InstalledProgram
	for rows.Next() {
		var app models.InstalledProgram
		if err := rows.Scan(
			&app.ClientID,
			&app.Name,
			&app.Version,
			&app.InstallDate,
			&app.Timestamp,
		); err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}

	return apps, nil
}

func (h *APIHandler) LatestDataHandler(w http.ResponseWriter, r *http.Request) {
	if h.ClientID == "" {
		http.Error(w, "Missing client_id parameter", http.StatusBadRequest)
		return
	}

	systemInfo, err := GetLatestSystemInfo(h.DB, h.ClientID)
	if err != nil {
		http.Error(w, "Failed to fetch system info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	apps, err := GetLatestInstalledApps(h.DB, h.ClientID)
	if err != nil {
		http.Error(w, "Failed to fetch apps: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		SystemInfo models.SystemInfo         `json:"system_info"`
		Apps       []models.InstalledProgram `json:"installed_apps"`
	}{
		SystemInfo: systemInfo,
		Apps:       apps,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
