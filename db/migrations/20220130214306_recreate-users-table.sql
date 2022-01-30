-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE users;

CREATE TABLE users (
    user_id     SERIAL NOT NULL,
    email       VARCHAR(100) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    role_id     INT NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id)    
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;

CREATE TABLE users (
    user_id     SERIAL NOT NULL,
    email       VARCHAR(100) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id)    
);
