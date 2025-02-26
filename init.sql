
create table if not exists users(
    "id" uuid not null primary key not null,
    "name" varchar(255) not null,
    "email" varchar(255) not null unique,
    "password" varchar(255) not null,
    "role" smallint not null
)

select * from users