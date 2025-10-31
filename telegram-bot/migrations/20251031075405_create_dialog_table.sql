-- +goose Up
-- +goose StatementBegin
CREATE TABLE dialog (
    id UUID PRIMARY KEY NOT NULL,
    step VARCHAR(255) NOT NULL,
    params JSONB NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT fk_dialog_user_id FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE RESTRICT
);
CREATE UNIQUE INDEX uk_dialog_user_id ON dialog (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE dialog;
-- +goose StatementEnd
