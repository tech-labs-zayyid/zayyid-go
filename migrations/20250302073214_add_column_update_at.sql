-- migrate:up
ALTER TABLE product_marketing.users 
ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- migrate:down
ALTER TABLE product_marketing.users 
DROP COLUMN updated_at;
