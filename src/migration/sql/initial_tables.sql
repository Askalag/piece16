create schema if not exists t1;

create table if not exists t1.tree
(
    id    bigserial,
    title varchar(100),
    tree_level smallint not null default 0
);

create table if not exists t1.task
(
    id         bigserial,
    title      varchar(100),
    tree_level smallint not null default 1,
    tree_id    int not null default 0
);

create table if not exists t1.task_item
(
    id         bigserial,
    title      varchar(100),
    tree_level smallint not null default 2,
    parent_id  int not null default 0
);

create table if not exists t1.time_item
(
    id          bigserial,
    title       varchar(100),
    description varchar(500),
    time_cost   decimal,
    tree_level  smallint not null default 3,
    parent_id   int not null default 0
);