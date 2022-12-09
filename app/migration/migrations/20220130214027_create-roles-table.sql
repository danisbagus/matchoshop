-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
    role_id  INTEGER NOT NULL,
    name        VARCHAR(50) NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (role_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE roles;
-- +goose StatementEnd


