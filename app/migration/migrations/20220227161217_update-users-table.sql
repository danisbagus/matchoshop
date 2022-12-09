-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN name VARCHAR(50) NULL;
-- +goose StatementEnd

-- +goose StatementBegin
UPDATE users SET name = 'Admin 1' WHERE user_id = 1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN name;
-- +goose StatementEnd
