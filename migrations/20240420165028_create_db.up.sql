CREATE TABLE IF NOT EXISTS orders
(
    order_uid          VARCHAR(255)                                                       PRIMARY KEY,
    track_number       VARCHAR(255)                                                       NOT NULL UNIQUE,
    entry              VARCHAR(255)                                                       NOT NULL,
    locale             VARCHAR(255)                                                       NOT NULL,
    internal_signature VARCHAR(255),
    customer_id        VARCHAR(255)                                                       NOT NULL,
    delivery_service   VARCHAR(255)                                                       NOT NULL,
    shardkey           VARCHAR(255)                                                       NOT NULL,
    sm_id              INT                                                                NOT NULL,
    date_created       TIMESTAMP                                                          NOT NULL DEFAULT NOW(),
    oof_shard          VARCHAR(255)                                                       NOT NULL
    );

CREATE TABLE IF NOT EXISTS  deliveries
(
    id             SERIAL       PRIMARY KEY,
    order_id       VARCHAR(255) REFERENCES orders (order_uid) ON DELETE CASCADE          NOT NULL,
    client_name    VARCHAR(255) NOT NULL,
    phone          VARCHAR(255) NOT NULL,
    zip            VARCHAR(255) NOT NULL,
    city           VARCHAR(255) NOT NULL,
    address        VARCHAR(255) NOT NULL,
    region         VARCHAR(255) NOT NULL,
    email          VARCHAR(255) NOT NULL
    );

CREATE TABLE IF NOT EXISTS  payments
(
    transaction_id VARCHAR(255) PRIMARY KEY,
    order_id       VARCHAR(255) REFERENCES orders (order_uid) ON DELETE CASCADE          NOT NULL,
    request_id     VARCHAR(255),
    currency       VARCHAR(255) NOT NULL,
    provider       VARCHAR(255) NOT NULL,
    amount         INT          NOT NULL,
    payment_dt     INT          NOT NULL,
    bank           VARCHAR(255) NOT NULL,
    delivery_cost  INT          NOT NULL,
    goods_total    INT          NOT NULL,
    custom_fee     INT          NOT NULL
    );

CREATE TABLE IF NOT EXISTS  items
(
    id           SERIAL       PRIMARY KEY,
    order_id     VARCHAR(255) REFERENCES orders (order_uid) ON DELETE CASCADE          NOT NULL,
    chrt_id      INT          NOT NULL,
    track_number VARCHAR(255) NOT NULL,
    price        INT          NOT NULL,
    rid          VARCHAR(255) NOT NULL,
    item_name    VARCHAR(255) NOT NULL,
    sale         INT          NOT NULL,
    item_size    VARCHAR(255) NOT NULL,
    total_price  INT          NOT NULL,
    nm_id        INT          NOT NULL,
    brand        VARCHAR(255) NOT NULL,
    status       INT          NOT NULL
    );