package data

import (
	"database/sql"
	"osqueryMVP/models"
)

func UpsertInstalledPrograms(db *sql.DB, programs []models.InstalledProgram) error {
	query := `
INSERT INTO installed_programs 
(client_id, name, version, install_date, timestamp)
VALUES (?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE
    version = VALUES(version),
    install_date = VALUES(install_date)
`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, p := range programs {
		_, err := stmt.Exec(p.ClientID, p.Name, p.Version, p.InstallDate, p.Timestamp)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpsertSystemInfo(db *sql.DB, info models.SystemInfo) error {
	query := `
REPLACE INTO system_info (
    client_id, hostname, cpu_brand, cpu_logical_cores, cpu_physical_cores,
    hardware_model, hardware_vendor, physical_memory, timestamp
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`
	_, err := db.Exec(query,
		info.ClientID,
		info.Hostname,
		info.CPUBrand,
		info.CPULogicalCores,
		info.CPUPhysicalCores,
		info.HardwareModel,
		info.HardwareVendor,
		info.PhysicalMemory,
		info.Timestamp,
	)
	return err
}
