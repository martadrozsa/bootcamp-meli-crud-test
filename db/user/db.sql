CREATE DATABASE IF NOT EXISTS `user`;

USE `user`;

DROP TABLE IF EXISTS `user`;

use user;

CREATE TABLE `user` (
    `id` MEDIUMINT NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `age` int(10) unsigned NOT NULL,
    `movie_genre` varchar(250) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table user;