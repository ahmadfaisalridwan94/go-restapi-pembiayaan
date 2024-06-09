BEGIN;
CREATE TABLE IF NOT EXISTS
    users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name varchar(255) NOT NULL,
        email varchar(255) NOT NULL UNIQUE,
        email_verify_at timestamp NULL,
        password varchar(255) NOT NULL,
        remember_token varchar(100) NULL,
        image varchar(255) NULL,
        role varchar(255) NOT NULL,
        status TINYINT NOT NULL DEFAULT 0,
        created_at timestamp NOT NULL,
        updated_at timestamp NOT NULL
    );
COMMIT;