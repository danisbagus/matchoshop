-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE products ADD COLUMN brand VARCHAR(50) NULL;

ALTER TABLE products ADD COLUMN image text NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE products DROP COLUMN brand;

ALTER TABLE products DROP COLUMN image;