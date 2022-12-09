-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    user_id     INTEGER NOT NULL,
    email       VARCHAR(100) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id)    
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE merchants (
    merchant_id INTEGER NOT NULL,
    user_id     INT NOT NULL,
    name        VARCHAR(50) NOT NULL,
    identifier  VARCHAR(50) NOT NULL UNIQUE,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (merchant_id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE products (
    product_id  INTEGER NOT NULL,
    name        VARCHAR(50) NOT NULL,
    sku         VARCHAR(20) NOT NULL,
    description VARCHAR(100) NULL,
    price       INT NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (product_id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE product_categories (
    product_category_id INTEGER NOT NULL,
    name                VARCHAR(50) NOT NULL,
    created_at          TIMESTAMP NOT NULL,
    updated_at          TIMESTAMP NOT NULL,
    PRIMARY KEY (product_category_id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE product_product_categories (
    product_id          INT NOT NULL,
    product_category_id INT NOT NULL,
    created_at          TIMESTAMP NOT NULL,
    updated_at          TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS merchants;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS product_categories;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS product_product_categories;
-- +goose StatementEnd