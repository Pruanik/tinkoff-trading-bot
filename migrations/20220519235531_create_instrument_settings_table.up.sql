create table if not exists instrument_settings
(
    id                           serial    primary key,
    figi                         varchar   not null,
	trading_strategy_settings_id integer   not null,
    is_data_collecting           boolean   not null default false,
	created_at                   timestamp not null,
	update_at                    timestamp not null
);

create index if not exists i_instrument_settings_figi
    on instrument_settings (figi);

create index if not exists i_instrument_settings_trading_strategy_settings_id
    on instrument_settings (trading_strategy_settings_id);

alter table instrument_settings 
add constraint fk_instrument_settings_instruments_figi 
foreign key (figi) 
references instruments (figi);

alter table instrument_settings 
add constraint fk_instrument_settings_trading_strategy_settings_id 
foreign key (trading_strategy_settings_id) 
references trading_strategy_settings (id);
