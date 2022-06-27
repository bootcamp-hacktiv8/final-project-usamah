begin;

create table if not exists social_medias(
    id serial primary key, 
	user_id int not null references users(id),
	name varchar(100) not null,
	social_media_url text not null,
	created_at timestamp not null,
    updated_at timestamp
);

commit;