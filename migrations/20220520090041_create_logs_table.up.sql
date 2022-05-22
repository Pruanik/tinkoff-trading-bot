create table if not exists logs
(
    id         bigserial primary key,
	category   varchar   not null,
	level      varchar   not null,
	message    varchar   not null,
	context    jsonb     not null,
	created_at timestamp not null
);

create index if not exists i_logs_category
    on logs (category);
create index if not exists i_logs_created_at
    on logs (created_at);
