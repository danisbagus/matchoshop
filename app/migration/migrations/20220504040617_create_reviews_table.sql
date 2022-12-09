-- +goose Up
-- +goose StatementBegin
CREATE TABLE reviews (
    review_id   INTEGER NOT NULL,
    user_id     INT NOT NULL,
    product_id  INT NOT NULL,
    rating      SMALLINT NOT NULL DEFAULT 0,
    comment     VARCHAR(50) NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (review_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reviews;
-- +goose StatementEnd


