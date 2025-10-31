-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
    id UUID PRIMARY KEY NOT NULL,
    chat_id BIGINT NOT NULL,
    first_name VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX uk_user_chat_id ON "user" (chat_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "user";
-- +goose StatementEnd
