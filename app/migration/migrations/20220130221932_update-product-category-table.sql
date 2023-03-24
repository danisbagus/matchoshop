-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE product_categories DROP COLUMN merchant_id;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE product_categories ADD COLUMN merchant_id INT NULL;

