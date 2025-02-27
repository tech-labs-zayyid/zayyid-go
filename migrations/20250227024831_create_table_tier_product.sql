-- migrate:up
CREATE TABLE
    IF NOT EXISTS product_marketing.product_tier (
    id VARCHAR(50) NOT NULL,
    tier_name VARCHAR(50) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id)
    );


-- migrate:down

