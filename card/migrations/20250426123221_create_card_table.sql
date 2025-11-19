-- +goose Up
-- +goose StatementBegin
CREATE TABLE card
(
    id         UUID         NOT NULL
        CONSTRAINT card_pk PRIMARY KEY,
    user_id    UUID         NOT NULL,
    question   VARCHAR(255) NOT NULL,
    answer     VARCHAR(255) NOT NULL,
    file_type  VARCHAR(255) NOT NULL,
    file_id    TEXT         NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE card;
-- +goose StatementEnd
