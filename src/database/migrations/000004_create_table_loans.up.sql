CREATE TABLE IF NOT EXISTS loans (
    id INT AUTO_INCREMENT PRIMARY KEY,
    borrower_id INT NOT NULL,
    lender_id INT NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    interest_rate DECIMAL(5,2) NOT NULL,
    term INT NOT NULL, -- in months
    status VARCHAR(50) NOT NULL, -- pending/approved/rejected
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    approved_at TIMESTAMP NULL,
    FOREIGN KEY (borrower_id) REFERENCES borrowers(id),
    FOREIGN KEY (lender_id) REFERENCES lenders(id)
);