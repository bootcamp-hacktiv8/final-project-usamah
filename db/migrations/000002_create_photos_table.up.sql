begin;

create table if not exists photos(
    id serial primary key, 
	user_id int not null references users(id),
	title varchar(100) not null,
	caption varchar(100),
    photo_url text not null,
    created_at timestamp not null,
    updated_at timestamp
);

commit;