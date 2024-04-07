-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE products DROP COLUMN merchant_id;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE products ADD COLUMN merchant_id INT NULL;

