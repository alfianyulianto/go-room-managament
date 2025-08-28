create table rooms(
    id bigint auto_increment,
    room_category_id bigint not null,
    code varchar(100) not null,
    name varchar(255) not null ,
    `condition` enum('Baik', 'Rusak Ringan', 'Rusak Sedang', 'Rusak Berat'),
    note text null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    primary key (id),
    foreign key (room_category_id) references room_categories(id) on update cascade on delete cascade,
    unique (code)
)engine = InnoDB;