-- +goose Up
-- +goose StatementBegin
CREATE TABLE review_log
(
    id         UUID                     NOT NULL
        CONSTRAINT card_review_pk PRIMARY KEY,
    card_id    UUID                     NOT NULL,
    rating     VARCHAR(255)             NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE review_log;
-- +goose StatementEnd