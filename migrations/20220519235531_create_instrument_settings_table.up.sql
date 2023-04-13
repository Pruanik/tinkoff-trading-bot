create table if not exists instrument_settings
(
    id                           serial    primary key,
    figi                         varchar   not null,
    is_data_collecting           boolean   not null default false,
	created_at                   timestamp not null,
	updated_at                   timestamp not null
);

create unique index if not exists i_instrument_settings_figi
    on instrument_settings (figi);

alter table instrument_settings 
add constraint fk_instrument_settings_instruments_figi 
foreign key (figi) 
references instruments (figi);
