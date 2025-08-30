create table room_categories(
    id bigint auto_increment,
    name varchar(255) not null ,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    primary key (id)
)engine = InnoDB;