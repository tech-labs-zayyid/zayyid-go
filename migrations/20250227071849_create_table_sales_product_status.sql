-- migrate:up
CREATE TABLE
    IF NOT EXISTS product_marketing.sales_product_status (
    id VARCHAR(50) NOT NULL,
    product_id VARCHAR(50) NOT NULL,
    status VARCHAR(30) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id)
    );

-- migrate:down
DROP TABLE product_marketing.sales_product_status
