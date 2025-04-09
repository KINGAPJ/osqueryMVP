CREATE TABLE IF NOT EXISTS installed_programs (
    client_id VARCHAR(64) NOT NULL,
    name VARCHAR(255) NOT NULL,
    version TEXT,
    install_date TEXT,
    timestamp DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (client_id, name)
);
