-- +goose Up
-- +goose StatementBegin

CREATE TABLE posts (
    id public.xid PRIMARY KEY DEFAULT xid(),
    title text,
    body text
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE posts;
-- +goose StatementEnd
