-- +goose Up
-- +goose StatementBegin
INSERT INTO payment_methods(payment_method_id, name) 
		  VALUES (1, 'Paypal');
-- +goose StatementEnd
	  
-- +goose Down
-- +goose StatementBegin
DELETE FROM payment_method_id WHERE role_id IN (1);
-- +goose StatementEnd