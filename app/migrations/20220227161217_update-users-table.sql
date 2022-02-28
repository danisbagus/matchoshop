-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users ADD COLUMN name VARCHAR(50) NULL;
UPDATE users SET name = 'Admin 1' WHERE user_id = 1;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE users DROP COLUMN name;

