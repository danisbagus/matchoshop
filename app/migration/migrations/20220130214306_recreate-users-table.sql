-- +goose Up
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

CREATE TABLE users (
    user_id     INTEGER NOT NULL,
    email       VARCHAR(100) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    role_id     INT NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id)    
);

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE users (
    user_id     INTEGER NOT NULL,
    email       VARCHAR(100) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id)    
);
-- +goose StatementEnd