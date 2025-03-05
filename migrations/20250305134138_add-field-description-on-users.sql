-- migrate:up
ALTER TABLE product_marketing.users 
ADD COLUMN description TEXT;

-- migrate:down
ALTER TABLE product_marketing.users 
DROP COLUMN description;
