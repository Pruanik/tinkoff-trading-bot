create table if not exists candles
(
    id         bigserial primary key,
	figi       varchar   not null,
	interval   integer,
	open       decimal   not null,
	high       decimal   not null,
	low        decimal   not null,
	close      decimal   not null,
	volume     bigserial not null,
	timestamp  timestamp not null,
	created_at timestamp not null
);

create index if not exists i_candles_figi
    on candles (figi);
create index if not exists i_candles_created_at
    on candles (created_at);
create index if not exists i_candles_timestamp
    on candles (timestamp);


alter table candles 
add constraint fk_candles_instruments_figi 
foreign key (figi) 
references instruments (figi);
