-- migrate:up
CREATE TABLE
    IF NOT EXISTS product_marketing.sales_product_status (
    id VARCHAR(50) NOT NULL,
    product_id VARCHAR(50) NOT NULL,
    status_id INT NOT NULL,
    status_name VARCHAR(30) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id)
    );

-- migrate:down

