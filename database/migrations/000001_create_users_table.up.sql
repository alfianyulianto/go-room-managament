create table users(
    id bigint auto_increment,
    name varchar(255) not null,
    email varchar(100) not null ,
    phone varchar(25) not null ,
    level enum('Admin', 'User') default 'User',
    password varchar(255) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    primary key (id)
)engine = InnoDB;