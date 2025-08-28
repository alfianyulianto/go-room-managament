create table room_reservations(
    id bigint auto_increment,
    room_id bigint not null ,
    user_id bigint not null ,
    start_at datetime not null ,
    end_at datetime not null ,
    purpose text,
    status enum('Pengajuan', 'Diterima', 'Ditlak'),
    approve_id bigint not null ,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    primary key (id),
    foreign key (room_id) references rooms(id) on update cascade on delete cascade,
    foreign key (user_id) references users(id) on update cascade on delete cascade,
    foreign key (approve_id) references users(id) on update cascade on delete cascade
)engine = InnoDB;