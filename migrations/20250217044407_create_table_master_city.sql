-- migrate:up
CREATE TABLE
    IF NOT EXISTS product_marketing.master_city (
    id VARCHAR(50) NOT NULL,
    province_id VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id)
    );


-- migrate:down

