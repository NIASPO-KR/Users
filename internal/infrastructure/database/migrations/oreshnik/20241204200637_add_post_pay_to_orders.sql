-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders
    ADD COLUMN postomat_id VARCHAR(36) NOT NULL DEFAULT '',
    ADD COLUMN payment_id VARCHAR(36) NOT NULL DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders
    DROP COLUMN postomat_id,
    DROP COLUMN payment_id;
-- +goose StatementEnd
