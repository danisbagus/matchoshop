-- +goose Up
-- +goose StatementBegin
ALTER TABLE products DROP COLUMN merchant_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE products ADD COLUMN merchant_id INT NULL;
-- +goose StatementEnd
