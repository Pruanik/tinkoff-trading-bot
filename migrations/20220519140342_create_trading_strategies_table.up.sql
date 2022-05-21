create table if not exists trading_strategies
(
    id 			serial    primary key,
	name 		varchar   not null,
	description text,
	created_at  timestamp not null
);
