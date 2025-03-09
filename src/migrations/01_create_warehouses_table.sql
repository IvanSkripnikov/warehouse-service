CREATE TABLE IF NOT EXISTS warehouses (
    id INT auto_increment PRIMARY KEY,
    title TEXT NOT NULL,
    volume INT NOT NULL,
    created BIGINT UNSIGNED,
    updated BIGINT UNSIGNED,
);