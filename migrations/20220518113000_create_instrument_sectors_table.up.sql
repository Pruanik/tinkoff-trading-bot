create table if not exists instrument_sectors
(
    id         serial          primary key,
	code       varchar         not null,
	name       varchar
);

create unique index if not exists u_instrument_sectors_code
    on instrument_sectors (code);