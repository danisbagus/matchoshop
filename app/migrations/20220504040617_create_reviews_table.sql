-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE reviews (
    review_id   SERIAL NOT NULL,
    user_id     INT NOT NULL,
    product_id  INT NOT NULL,
    rating      SMALLINT NOT NULL DEFAULT 0,
    comment     VARCHAR(50) NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (review_id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE reviews;


