package models

import "time"

type InstalledProgram struct {
	ClientID    string    `db:"client_id"`
	Name        string    `db:"name"`
	Version     string    `db:"version"`
	InstallDate string    `db:"install_date"`
	Timestamp   time.Time `db:"timestamp"`
}

type SystemInfo struct {
	ClientID         string    `db:"client_id"`
	Hostname         string    `db:"hostname"`
	CPUBrand         string    `db:"cpu_brand"`
	CPULogicalCores  int       `db:"cpu_logical_cores"`
	CPUPhysicalCores int       `db:"cpu_physical_cores"`
	HardwareModel    string    `db:"hardware_model"`
	HardwareVendor   string    `db:"hardware_vendor"`
	PhysicalMemory   string    `db:"physical_memory"`
	Timestamp        time.Time `db:"timestamp"`
}
