CREATE TABLE IF NOT EXISTS borrowers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    address TEXT NOT NULL,
    date_of_birth DATE NOT NULL,
    id_number VARCHAR(50) NOT NULL,
    tax_number VARCHAR(50),
    FOREIGN KEY (user_id) REFERENCES users(id)
);