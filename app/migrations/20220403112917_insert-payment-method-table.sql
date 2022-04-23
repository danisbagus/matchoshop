-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO payment_methods(payment_method_id, name) 
		  VALUES (1, 'Paypal');
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM roles WHERE role_id IN (1,2,3,4);