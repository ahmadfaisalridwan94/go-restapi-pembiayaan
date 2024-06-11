CREATE TABLE IF NOT EXISTS loan_products (
    loan_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    PRIMARY KEY (loan_id, product_id),
    FOREIGN KEY (loan_id) REFERENCES loans(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);