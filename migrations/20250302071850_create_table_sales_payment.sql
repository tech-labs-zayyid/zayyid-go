-- migrate:up
CREATE TABLE
    IF NOT EXISTS product_marketing.sales_payment (
    transaction_id VARCHAR(50) NOT NULL,
    transaction_status VARCHAR(30) NOT NULL DEFAULT "",
    status_message VARCHAR(150) NOT NULL DEFAULT "",
    status_code INTEGER NOT NULL DEFAULT 0,
    payment_type VARCHAR(20) NOT NULL DEFAULT "",
    order_id VARCHAR(20) NOT NULL,
    gross_amount NUMERIC(15,2) NOT NULL DEFAULT 0,
    fraud_status VARCHAR(15) NOT NULL DEFAULT "",
    bank VARCHAR(10) NOT NULL DEFAULT "",
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (transaction_id, order_id)
    );

-- migrate:down
DROP TABLE product_marketing.sales_payment
