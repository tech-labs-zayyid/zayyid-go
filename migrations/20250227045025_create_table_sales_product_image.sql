-- migrate:up
CREATE TABLE
    IF NOT EXISTS product_marketing.sales_product_image (
    id VARCHAR(50) NOT NULL,
    product_id VARCHAR(50) NOT NULL,
    image_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id)
    );


-- migrate:down
DROP TABLE product_marketing.sales_product_description
