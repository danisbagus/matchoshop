-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE merchants;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
CREATE TABLE merchants (
    merchant_id SERIAL NOT NULL,
    user_id     INT NOT NULL,
    name        VARCHAR(50) NOT NULL,
    identifier  VARCHAR(50) NOT NULL UNIQUE,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (merchant_id)
);