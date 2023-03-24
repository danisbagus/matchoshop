-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE product_product_categories
DROP COLUMN created_at,
DROP COLUMN updated_at;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE product_product_categories
ADD COLUMN created_at TIMESTAMP NOT NULL,
ADD COLUMN updated_at TIMESTAMP NOT NULL;