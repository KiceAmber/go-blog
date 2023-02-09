create database `go_blog`;

use `go_blog`;

create table `user` (
    `id` bigint(20) not null auto_increment primary key,
    `username` varchar(64) not null,
    `password` varchar(20) not null,
    unique index `idx_username`(`username`)
) engine=innodb;