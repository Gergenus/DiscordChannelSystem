-- +goose Up
-- +goose StatementBegin
CREATE TABLE channels (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_by INT REFERENCES users(uid) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE channels;
-- +goose StatementEnd
