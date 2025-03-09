CREATE TABLE IF NOT EXISTS warehouse_items (
    id INT auto_increment PRIMARY KEY,
    warehouse_id INT NOT NULL,
    item_id INT NOT NULL,
    volume INT NOT NULL,
    created BIGINT UNSIGNED,
    updated BIGINT UNSIGNED,
    status TINYINT DEFAULT 1
);