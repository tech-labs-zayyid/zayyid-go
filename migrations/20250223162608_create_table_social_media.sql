-- migrate:up
CREATE TABLE
    IF NOT EXISTS product_marketing.sales_social_media (
    id VARCHAR(50) NOT NULL,
    sales_id VARCHAR(50) NOT NULL,
    public_access VARCHAR(20) NOT NULL,
    social_media_name VARCHAR(30) NOT NULL,
    user_account VARCHAR(30) NOT NULL,
    link_embed text,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id)
    );


-- migrate:down

