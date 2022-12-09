-- +goose Up
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE product_categories;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE products (
    product_id  INTEGER NOT NULL,
    merchant_id INT NOT NULL,
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
    merchant_id         INT NOT NULL,
    name                VARCHAR(50) NOT NULL,
    created_at          TIMESTAMP NOT NULL,
    updated_at          TIMESTAMP NOT NULL,
    PRIMARY KEY (product_category_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE product_categories;
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
