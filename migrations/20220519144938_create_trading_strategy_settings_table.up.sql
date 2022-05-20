create table if not exists trading_strategy_settings
(
    id                  integer   primary key,
	trading_strategy_id integer   not null,
	notes 	            text,
	parameters          jsonb     not null,
	created_at          timestamp not null
);

create index if not exists i_trading_strategy_settings_trading_strategy_id
    on trading_strategy_settings (trading_strategy_id);

alter table trading_strategy_settings 
add constraint fk_trading_strategy_settings_trading_strategy_id 
foreign key (trading_strategy_id) 
references trading_strategies (id);
