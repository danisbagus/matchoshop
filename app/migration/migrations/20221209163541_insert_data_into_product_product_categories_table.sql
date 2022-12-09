-- +goose Up
-- +goose StatementBegin
INSERT INTO product_product_categories(product_id, product_category_id) 
		  VALUES (1, 1),
                 (2, 1),
                 (3, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM product_product_categories WHERE product_id IN (1,2,3);
-- +goose StatementEnd