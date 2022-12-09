-- +goose Up
-- +goose StatementBegin
ALTER TABLE product_categories DROP COLUMN merchant_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE product_categories ADD COLUMN merchant_id INT NULL;
-- +goose StatementEnd
