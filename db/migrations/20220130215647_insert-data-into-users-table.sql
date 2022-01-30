-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO users(user_id, email, password, role_id, created_at, updated_at) 
		   VALUES(1, 'matchoshop@live.com', '$2a$14$rAfbQlhIk38nfIZ8aOIOueDhYpq4hwetQXjwpvMuNMlY0Q1PoveyG', 1,current_timestamp, current_timestamp);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM users WHERE user_id = 1;