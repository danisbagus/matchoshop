-- +goose Up
-- +goose StatementBegin
DROP TABLE merchants;
-- +goose StatementEnd

-- +goose Down
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
