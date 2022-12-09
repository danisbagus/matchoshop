-- +goose Up
-- +goose StatementBegin
INSERT INTO product_categories(product_category_id, name, created_at, updated_at) 
		  VALUES (1, 'Tas', current_timestamp, current_timestamp);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM product_categories WHERE product_category_id IN (1);
-- +goose StatementEnd