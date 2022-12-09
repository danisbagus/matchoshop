-- +goose Up
-- +goose StatementBegin
CREATE TABLE refresh_token_stores (
    refresh_token   VARCHAR(300) NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    PRIMARY KEY (refresh_token)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE refresh_token_stores;
-- +goose StatementEnd


