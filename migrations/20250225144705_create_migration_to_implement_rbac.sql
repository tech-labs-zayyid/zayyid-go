-- migrate:up
CREATE TYPE sales.user_role AS ENUM ('sales', 'agent'); -- Buat ENUM dulu

CREATE TABLE 
    IF NOT EXISTS sales.users (
    id VARCHAR(36) PRIMARY KEY,
    username VARCHAR(50) UNIQUE,
    name VARCHAR(255),
    whatsapp_number VARCHAR(13),
    email VARCHAR(100),
    password VARCHAR(255),
    role sales.user_role NOT NULL, -- Gunakan ENUM yang sudah dibuat
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(36)
);

CREATE TABLE 
    IF NOT EXISTS sales.sales_agent (
    sales_id VARCHAR(36),
    agent_id VARCHAR(36),
    PRIMARY KEY (sales_id, agent_id),
    FOREIGN KEY (sales_id) REFERENCES sales.users(id) ON DELETE CASCADE,
    FOREIGN KEY (agent_id) REFERENCES sales.users(id) ON DELETE CASCADE
);

-- migrate:down
DROP TABLE sales.sales_agent;
DROP TABLE sales.users;
DROP TYPE sales.user_role; -- Hapus ENUM saat rollback
