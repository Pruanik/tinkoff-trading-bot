create type instrument_type as enum ('bond', 'currency', 'etf', 'future', 'share');

create table if not exists instruments
(
    id integer primary key,
	figi varchar not null,
	type instrument_type not null,
	created_at timestamp not null,
);

create index if not exists i_instruments_figi
    on instruments (figi);
