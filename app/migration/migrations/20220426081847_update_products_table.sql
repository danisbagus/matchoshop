-- +goose Up
-- +goose StatementBegin
ALTER TABLE products ADD COLUMN brand VARCHAR(50) NULL;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE products ADD COLUMN image text NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE products DROP COLUMN brand;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE products DROP COLUMN image;
-- +goose StatementEnd
