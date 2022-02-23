-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE roles (
    role_id  SERIAL NOT NULL,
    name        VARCHAR(50) NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (role_id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE roles;


