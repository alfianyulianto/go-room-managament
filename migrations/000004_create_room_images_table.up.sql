create table room_images(
    id bigint auto_increment,
    room_id bigint not null ,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    primary key (id),
    foreign key (room_id)references rooms(id)
)engine = InnoDB;