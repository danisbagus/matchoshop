-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- order table
CREATE TABLE orders (
    order_id            SERIAL NOT NULL,
    user_id             INT NOT NULL,
    payment_method_id   INT NOT NULL,
    product_price       INT NOT NULL,
    tax_price           INT NOT NULL,
    shipping_price      INT NOT NULL,
    total_price         INT NOT NULL,
    is_paid             SMALLINT NOT NULL DEFAULT 0,
    paid_at             TIMESTAMP NULL,
    is_delivered        SMALLINT NOT NULL DEFAULT 0,
    delivered_at        TIMESTAMP NULL,
    created_at          TIMESTAMP NOT NULL,
    updated_at          TIMESTAMP NOT NULL,
    PRIMARY KEY (order_id)    
);

-- payment method table
CREATE TABLE payment_methods (
    payment_method_id   SERIAL NOT NULL,
    name                VARCHAR(50) NOT NULL,
    PRIMARY KEY (payment_method_id)    
);

-- order product category table
CREATE TABLE order_products (
    order_id            INT NOT NULL,
    product_id          INT NOT NULL,
    quantity            INT NOT NULL
);

-- payment result table
CREATE TABLE payment_results (
    payment_result_id   VARCHAR(20) NOT NULL,
    order_id            INT NOT NULL,
    status              VARCHAR(10) NULL,
    update_time         TIMESTAMP NOT NULL,
    email               VARCHAR(100) NOT NULL,
    PRIMARY KEY (payment_result_id)
);

-- shipment address table
CREATE TABLE shipment_address (
    shipment_address_id   SERIAL NOT NULL,
    order_id              INT NOT NULL,
    address               VARCHAR(100) NOT NULL,
    city                  VARCHAR(20) NOT NULL,
    postal_code           VARCHAR(10) NOT NULL,
    country               VARCHAR(20) NOT NULL,
    PRIMARY KEY (shipment_address_id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE IF EXISTS orders;

DROP TABLE IF EXISTS payment_methods;

DROP TABLE IF EXISTS order_products;

DROP TABLE IF EXISTS payment_results;

DROP TABLE IF EXISTS shipment_address;