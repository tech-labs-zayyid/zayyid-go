-- migrate:up
ALTER TABLE product_marketing.sales_agent 
DROP CONSTRAINT IF EXISTS sales_agent_sales_id_fkey;

ALTER TABLE product_marketing.sales_agent 
ADD CONSTRAINT sales_agent_sales_id_fkey
FOREIGN KEY (sales_id) REFERENCES product_marketing.users(id)
ON DELETE CASCADE ON UPDATE CASCADE;

-- migrate:down
ALTER TABLE product_marketing.sales_agent 
DROP CONSTRAINT IF EXISTS sales_agent_sales_id_fkey;
