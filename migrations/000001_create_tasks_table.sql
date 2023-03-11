-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT exists notifications
(
    id          serial PRIMARY KEY,
    title       varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    user_id     serial       NOT NULL,
    created_at  timestamp(0) NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notifications;
-- +goose StatementEnd
