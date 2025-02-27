-- migrate:up
CREATE TABLE
    IF NOT EXISTS product_marketing.product_tier_detail (
    id VARCHAR(50) NOT NULL,
    tier_id VARCHAR(50) NOT NULL,
    feature VARCHAR(50) NOT NULL,
    limitation VARCHAR(50) NOT NULL,
    length_limitation VARCHAR(50) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id)
    );


-- migrate:down

