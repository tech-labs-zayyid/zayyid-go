-- migrate:up
ALTER TABLE product_marketing.sales_agent 
ADD COLUMN created_at TIMESTAMP NOT NULL,
ADD COLUMN created_by VARCHAR(36) NOT NULL;

-- This migration below already implemented using manual method
-- ALTER TABLE product_marketing.users 
-- ADD COLUMN description TEXT;

-- migrate:down
ALTER TABLE product_marketing.sales_agent 
DROP COLUMN created_at,
DROP COLUMN created_by;

-- This migration below already implemented using manual method
-- ALTER TABLE product_marketing.users 
-- DROP COLUMN description;
