CREATE TABLE IF NOT EXISTS lenders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    company_name VARCHAR(100),
    address TEXT,
    license_number VARCHAR(50),
    FOREIGN KEY (user_id) REFERENCES users(id)
);