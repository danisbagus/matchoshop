-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE refresh_token_stores (
    refresh_token   VARCHAR(300) NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    PRIMARY KEY (refresh_token)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE refresh_token_stores;


