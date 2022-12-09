-- +goose Up
-- +goose StatementBegin
INSERT INTO products(product_id, name, sku, brand, image, description, price, stock, created_at, updated_at) 
		  VALUES (1, 'Tas Ransel WC001', 'SKU1', 'Panamino', 'https://res.cloudinary.com/matchoshop/image/upload/v1670577652/matchoshop/tas_ransel_1_jxmo0i.jpg', 'Panamino bag new model', 500000, 10, current_timestamp, current_timestamp),
                 (2, 'Tas Ransel WC002', 'SKU2', 'Panamino', 'https://res.cloudinary.com/matchoshop/image/upload/v1670577652/matchoshop/tas_ransel_2_dh5tan.jpg', 'Panamino bag new model', 500000, 10, current_timestamp, current_timestamp),
                 (3, 'Tas Ransel WC003', 'SKU3', 'Panamino', 'https://res.cloudinary.com/matchoshop/image/upload/v1670577652/matchoshop/tas_ransel_3_hew3ef.jpg', 'Panamino bag new model', 500000, 10, current_timestamp, current_timestamp);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM products WHERE product_id IN (1,2,3);
-- +goose StatementEnd