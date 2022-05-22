create table if not exists currencies
(
    id                        serial    primary key,
	figi                      varchar   not null,
	ticker                    varchar   not null,
	class_code                varchar   not null,
	isin                      varchar   not null,
	lot                       integer   not null,
	currency                  varchar   not null,
	name                      varchar   not null,
	exchange                  varchar   not null,
    otc_flag                  boolean   not null,
    buy_available_flag        boolean   not null,
    sell_available_flag       boolean   not null,
    iso_currency_name         varchar   not null,
	min_price_increment_units bigint    not null,
	min_price_increment_nano  integer   not null,
	api_trade_available_flag  boolean   not null,
	created_at                timestamp not null,
	updated_at                timestamp not null
);

create unique index if not exists i_currencies_figi
    on currencies (figi);

alter table currencies 
add constraint fk_currencies_instruments_figi 
foreign key (figi) 
references instruments (figi);
