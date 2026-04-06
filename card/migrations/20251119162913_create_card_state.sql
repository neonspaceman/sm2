-- +goose Up
-- +goose StatementBegin
CREATE TABLE card_state
(
    id                       UUID                     NOT NULL
        CONSTRAINT card_state_pk PRIMARY KEY,
    state                    VARCHAR(255)             NOT NULL,
    step                     INTEGER                  NOT NULL,
    easy                     DOUBLE PRECISION         NOT NULL,
    due                      TIMESTAMP WITH TIME ZONE NOT NULL,
    current_interval_in_days BIGINT                   NOT NULL,
    created_at               TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at               TIMESTAMP WITH TIME ZONE NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE card_state;
-- +goose StatementEnd
