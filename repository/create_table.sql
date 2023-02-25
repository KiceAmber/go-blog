CREATE DATABASE `go_blog`;

USE `go_blog`;

-- ---------------
--   用户信息表
-- ---------------
CREATE TABLE `user` (
        `id` BIGINT(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `user_id` BIGINT(20) NOT NULL,
        `username` VARCHAR(32) NOT NULL UNIQUE,
        `password` VARCHAR(64) NOT NULL,
        `gender` TINYINT(2) NOT NULL DEFAULT 0,
        `email` VARCHAR(30),
        `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
        `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        UNIQUE INDEX `idx_user_id` (`user_id`),
        UNIQUE INDEX `idx_username` (`username`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 collate=utf8mb4_general_ci;


-- ---------------
--   文章表
-- ---------------
CREATE TABLE `article` (
        `id` BIGINT(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `article_id` BIGINT(20) NOT NULL,
        `title` VARCHAR(64) NOT NULL,
        `content` LONGTEXT NOT NULL,
        `cover_image` VARCHAR(64) NOT NULL,
        `author_id` BIGINT(20) NOT NULL,
        `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
        `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        INDEX `idx_title` (`title`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 collate=utf8mb4_general_ci;