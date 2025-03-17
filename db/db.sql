CREATE DATABASE IF NOT EXISTS retailer_db;
USE retailer_db;
CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    status ENUM('order placed', 'processed', 'failed') DEFAULT 'order placed',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);
SHOW TABLES;
DESCRIBE products;
DESCRIBE orders;
ALTER TABLE orders MODIFY status ENUM('order placed', 'processed', 'failed') DEFAULT 'order placed' NOT NULL;
DESCRIBE orders;
UPDATE orders SET status = 'order placed' WHERE status IS NULL;
select * from products;
select * from orders;