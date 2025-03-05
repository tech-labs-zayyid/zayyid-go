-- migrate:up
ALTER TABLE product_marketing.sales_agent DROP CONSTRAINT sales_agent_agent_id_fkey;

-- migrate:down
ALTER TABLE product_marketing.sales_agent ADD CONSTRAINT sales_agent_agent_id_fkey
FOREIGN KEY (agent_id) REFERENCES product_marketing.users(id);
