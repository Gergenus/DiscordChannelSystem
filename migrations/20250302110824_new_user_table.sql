-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    uid SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    hashpassword TEXT NOT NULL 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd
