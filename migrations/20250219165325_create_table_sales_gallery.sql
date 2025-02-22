-- migrate:up
CREATE TABLE
    IF NOT EXISTS product_marketing.sales_gallery (
    id VARCHAR(50) NOT NULL,
    sales_id VARCHAR(50) NOT NULL,
    public_access VARCHAR(20) NOT NULL,
    image_url text NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id)
    );


-- migrate:down

