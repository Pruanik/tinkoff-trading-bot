create table if not exists shares
(
    id integer primary key,
	figi varchar not null,
	ticker varchar not null,
	class_code varchar not null,
	isin varchar not null,
	lot integer not null,
	currency varchar not null,
	name varchar not null,
	exchange varchar not null,
	sector varchar not null,
	min_price_increment_units bigint not null,
	min_price_increment_nano integer not null,
	api_trade_available_flag boolean not null,
	created_at timestamp not null,
	updated_at timestamp not null,
);

create index if not exists i_shares_figi
    on shares (figi);

alter table shares 
add constraint fk_shares_instruments_figi 
foreign key (figi) 
references instruments (figi);
