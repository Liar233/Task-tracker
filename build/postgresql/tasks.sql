create table tasks
(
    name   text         not null,
    owner  varchar(255) not null,
    status integer      not null
);

alter table tasks
    owner to tracker;

