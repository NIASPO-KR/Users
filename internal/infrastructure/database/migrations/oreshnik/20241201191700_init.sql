-- +goose Up
-- +goose StatementBegin
CREATE TABLE carts
(
    user_id VARCHAR(36) NOT NULL DEFAULT 'GOLBUTSA-1337-1487-911Z-Salla4VO2022',
    item_id VARCHAR(36) NOT NULL,
    count   INTEGER     NOT NULL DEFAULT 1
        CONSTRAINT check_positive_count CHECK (count > 0)
);
CREATE INDEX idx_cart_user_id ON carts(user_id);
CREATE INDEX idx_cart_item_id ON carts(item_id);

CREATE TYPE status_type AS ENUM ('В работе', 'Доставляется', 'Получен', 'Отменен', 'Отказ');
CREATE TABLE orders
(
    id      VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL DEFAULT 'GOLBUTSA-1337-1487-911Z-Salla4VO2022',
    status  status_type NOT NULL DEFAULT 'В работе'
);
CREATE INDEX idx_orders_id ON orders(id);
CREATE INDEX idx_orders_user_id ON orders(user_id);

CREATE TABLE orders_items
(
    order_id VARCHAR(36) NOT NULL REFERENCES orders (id) ON DELETE CASCADE,
    item_id  VARCHAR(36) NOT NULL,
    count    INTEGER     NOT NULL
        CONSTRAINT check_positive_count CHECK (count > 0)
);
CREATE INDEX idx_orders_items_order_id ON orders_items(order_id);
CREATE INDEX idx_orders_items_item_id ON orders_items(item_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS orders_items;

DROP INDEX IF EXISTS idx_cart_user_id;
DROP INDEX IF EXISTS idx_cart_item_id;
DROP INDEX IF EXISTS idx_orders_id;
DROP INDEX IF EXISTS idx_orders_user_id;
DROP INDEX IF EXISTS idx_orders_items_order_id;
DROP INDEX IF EXISTS idx_orders_items_item_id;
-- +goose StatementEnd