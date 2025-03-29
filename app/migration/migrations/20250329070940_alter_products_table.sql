-- +goose Up
-- +goose StatementBegin
ALTER TABLE products ALTER COLUMN description TYPE VARCHAR(500);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin 
ALTER TABLE products ALTER COLUMN description TYPE VARCHAR(100);
-- +goose StatementEnd
