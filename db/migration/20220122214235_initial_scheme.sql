-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- users table
CREATE TABLE users (
    user_id     SERIAL NOT NULL,
    email       VARCHAR(100) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id)    
);

-- merchants table
CREATE TABLE merchants (
    merchant_id SERIAL NOT NULL,
    user_id     INT NOT NULL,
    name        VARCHAR(50) NOT NULL,
    identifier  VARCHAR(50) NOT NULL UNIQUE,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (merchant_id)
);

-- products table
CREATE TABLE products (
    product_id  SERIAL NOT NULL,
    name        VARCHAR(50) NOT NULL,
    sku         VARCHAR(20) NOT NULL,
    description VARCHAR(100) NULL,
    price       INT NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (product_id)
);

-- product category table
CREATE TABLE product_categories (
    product_category_id SERIAL NOT NULL,
    name                VARCHAR(50) NOT NULL,
    created_at          TIMESTAMP NOT NULL,
    updated_at          TIMESTAMP NOT NULL,
    PRIMARY KEY (product_category_id)
);

-- product product category table
CREATE TABLE product_product_categories (
    product_id          INT NOT NULL,
    product_category_id INT NOT NULL,
    created_at          TIMESTAMP NOT NULL,
    updated_at          TIMESTAMP NOT NULL
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS merchants;

DROP TABLE IF EXISTS products;

DROP TABLE IF EXISTS product_categories;

DROP TABLE IF EXISTS product_product_categories;