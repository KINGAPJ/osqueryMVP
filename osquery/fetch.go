package osquery

import (
	"encoding/json"
	"os/exec"
	"strconv"
	"time"

	"osqueryMVP/models"
)

func GetClientID() (string, error) {
	out, err := exec.Command("osqueryi.exe", "--json", "SELECT uuid FROM system_info;").Output()
	if err != nil {
		return "", err
	}
	var res []map[string]string
	if err := json.Unmarshal(out, &res); err != nil {
		return "", err
	}
	return res[0]["uuid"], nil
}

func FetchInstalledPrograms(clientID string) ([]models.InstalledProgram, error) {
	query := `SELECT name, version, install_source, install_date FROM programs WHERE name IS NOT NULL LIMIT 10;`
	out, err := runOsquery(query)
	if err != nil {
		return nil, err
	}
	var rows []struct {
		Name          string `json:"name"`
		Version       string `json:"version"`
		InstallSource string `json:"install_source"`
		InstallDate   string `json:"install_date"`
	}
	err = json.Unmarshal(out, &rows)
	if err != nil {
		return nil, err
	}
	var programs []models.InstalledProgram
	for _, r := range rows {
		programs = append(programs, models.InstalledProgram{
			ClientID:    clientID,
			Name:        r.Name,
			Version:     r.Version,
			InstallDate: r.InstallDate,
			Timestamp:   time.Now(),
		})
	}
	return programs, nil
}

func FetchSystemInfo(clientID string) (models.SystemInfo, error) {
	query := `SELECT * FROM system_info LIMIT 10;`
	out, err := runOsquery(query)
	if err != nil {
		return models.SystemInfo{}, err
	}
	var rows []map[string]interface{}
	if err := json.Unmarshal(out, &rows); err != nil {
		return models.SystemInfo{}, err
	}
	row := rows[0]
	cpuLogicalCores, err := strconv.Atoi(row["cpu_logical_cores"].(string))
	if err != nil {
		return models.SystemInfo{}, err
	}
	cpuPhysicalCores, err := strconv.Atoi(row["cpu_physical_cores"].(string))
	if err != nil {
		return models.SystemInfo{}, err
	}
	return models.SystemInfo{
		ClientID:         clientID,
		Hostname:         row["hostname"].(string),
		CPUBrand:         row["cpu_brand"].(string),
		CPULogicalCores:  cpuLogicalCores,
		CPUPhysicalCores: cpuPhysicalCores,
		HardwareModel:    row["hardware_model"].(string),
		HardwareVendor:   row["hardware_vendor"].(string),
		PhysicalMemory:   row["physical_memory"].(string),
		Timestamp:        time.Now(),
	}, nil
}

func runOsquery(query string) ([]byte, error) {
	cmd := exec.Command("osqueryi.exe", "--json", query)
	return cmd.Output()
}
