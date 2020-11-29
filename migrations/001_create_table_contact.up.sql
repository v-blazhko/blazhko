create table contact
(
    id          int auto_increment primary key,
    client_name varchar(128)  null,
    email       varchar(128)  null,
    message     varchar(1000) null,
    created_at  timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);