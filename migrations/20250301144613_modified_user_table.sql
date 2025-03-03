-- migrate:up
ALTER TABLE product_marketing.users 
ADD COLUMN image_url VARCHAR(255),
ADD COLUMN referal_code VARCHAR(50);

CREATE TABLE product_marketing.agent_social_media (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36),
    platform_type VARCHAR(20),
    value VARCHAR(50),
    created_at TIMESTAMP
);

-- migrate:down
ALTER TABLE product_marketing.users 
DROP COLUMN image_url,
DROP COLUMN referal_code;
DROP TABLE product_marketing.agent_social_media;
