create table if not exists tasks
(
    id             bigserial   primary key,
	label          varchar     not null,
    arguments      jsonb       not null,
    status         smallserial not null,
    is_rescheduled boolean     not null,
    time_shift     serial,
	executed_at    timestamp   not null,
	created_at     timestamp   not null,
	updated_at     timestamp   not null
);

create index if not exists i_tasks_status
    on tasks (status);
create index if not exists i_tasks_executed_at
    on tasks (executed_at);
