-- +goose Up
-- +goose StatementBegin
ALTER TABLE product_product_categories DROP COLUMN created_at;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE product_product_categories DROP COLUMN updated_at;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE product_product_categories
ADD COLUMN created_at TIMESTAMP NOT NULL,
ADD COLUMN updated_at TIMESTAMP NOT NULL;
-- +goose StatementEnd
