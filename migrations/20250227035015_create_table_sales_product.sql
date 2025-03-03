-- migrate:up
CREATE TABLE
    IF NOT EXISTS product_marketing.sales_product (
    id VARCHAR(50) NOT NULL,
    sales_id varchar(50) NOT NULL,
    public_access VARCHAR(20) NOT NULL,
    slug TEXT NOT NULL,
    page_category_name VARCHAR(50) NOT NULL,
    sub_category_product VARCHAR(50),
    product_name TEXT NOT NULL,
    price NUMERIC DEFAULT 0,
    tdp NUMERIC DEFAULT 0,
    installment NUMERIC DEFAULT 0,
    best_product BOOLEAN NOT NULL DEFAULT FALSE,
    city_id VARCHAR(50) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id)
    );


-- migrate:down
DROP TABLE product_marketing.sales_product

