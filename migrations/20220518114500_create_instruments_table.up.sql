drop type if exists instrument_type;
create type instrument_type as enum ('bond', 'currency', 'etf', 'future', 'share');

create table if not exists instruments
(
    id         serial          primary key,
	name       varchar         not null,
	figi       varchar         not null,
	sector_id  serial          not null,
	type       instrument_type not null,
	created_at timestamp       not null
);

create unique index if not exists i_instruments_figi
    on instruments (figi);

alter table instruments 
add constraint fk_instruments_instrument_sectors_sector_id 
foreign key (sector_id) 
references instrument_sectors (id);
