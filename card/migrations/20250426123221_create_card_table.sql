-- +goose Up
-- +goose StatementBegin
CREATE TABLE card
(
    id            UUID         NOT NULL
        CONSTRAINT card_pk PRIMARY KEY,
    front_content VARCHAR(255) NOT NULL,
    back_content  TEXT         NOT NULL,
    created_at    TIMESTAMP    NOT NULL,
    updated_at    TIMESTAMP    NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE card;
-- +goose StatementEnd
