create table task (
    id bigserial,
    title varchar(100),
    tree_level bytea
);

create table task_item (
    id bigserial,
    title varchar(100),
    tree_level bytea,
    parent_id int
);

create table time_item (
    id bigserial,
    title varchar(100),
    description varchar(500),
    time_cost decimal,
    tree_level bytea,
    parent_id int
);