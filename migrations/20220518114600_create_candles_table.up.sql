create table if not exists candles
(
    id         bigserial primary key,
    created_at timestamp,
    updated_at timestamp
);
