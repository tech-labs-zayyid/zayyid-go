-- migrate:up
CREATE TABLE
    IF NOT EXISTS product_marketing.sales_testimony (
	id VARCHAR(50) NOT NULL,
    public_access VARCHAR(20) NOT NULL,
	fullname VARCHAR(150) NOT NULL DEFAULT '',
	description TEXT NOT NULL,
	photo_url TEXT NOT NULL,
	is_active BOOLEAN NOT NULL DEFAULT TRUE,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT NULL,
	PRIMARY KEY (id)
);

-- migrate:down
