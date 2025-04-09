CREATE TABLE IF NOT EXISTS system_info (
    client_id VARCHAR(64) NOT NULL PRIMARY KEY,
    hostname TEXT,
    cpu_brand TEXT,
    cpu_logical_cores INTEGER,
    cpu_physical_cores INTEGER,
    hardware_model TEXT,
    hardware_vendor TEXT,
    physical_memory TEXT,
    timestamp DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
