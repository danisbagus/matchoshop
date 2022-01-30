-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO roles(role_id, name, created_at, updated_at) 
		  VALUES (1, 'Super Admin', current_timestamp, current_timestamp),
                 (2, 'Admin', current_timestamp, current_timestamp),
                 (3, 'Customer', current_timestamp, current_timestamp),
                 (4, 'Guest', current_timestamp, current_timestamp);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM roles WHERE role_id IN (1,2,3,4);