begin;

create table if not exists users(
    id serial primary key, 
	username varchar(100) not null unique,
	email varchar(100) not null unique,
	password varchar(100) not null,
    age int not null,
    created_at timestamp not null,
    updated_at timestamp
);

commit;