-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE products;
DROP TABLE product_categories;

CREATE TABLE products (
    product_id  SERIAL NOT NULL,
    merchant_id INT NOT NULL,
    name        VARCHAR(50) NOT NULL,
    sku         VARCHAR(20) NOT NULL,
    description VARCHAR(100) NULL,
    price       INT NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (product_id)
);

CREATE TABLE product_categories (
    product_category_id SERIAL NOT NULL,
    merchant_id         INT NOT NULL,
    name                VARCHAR(50) NOT NULL,
    created_at          TIMESTAMP NOT NULL,
    updated_at          TIMESTAMP NOT NULL,
    PRIMARY KEY (product_category_id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE products;
DROP TABLE product_categories;

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

CREATE TABLE product_categories (
    product_category_id SERIAL NOT NULL,
    name                VARCHAR(50) NOT NULL,
    created_at          TIMESTAMP NOT NULL,
    updated_at          TIMESTAMP NOT NULL,
    PRIMARY KEY (product_category_id)
);
