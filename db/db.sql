CREATE DATABASE IF NOT EXISTS `movie`;

USE `movie`;

DROP TABLE IF EXISTS `movie`;

use movie;

CREATE TABLE `movie` (
`id` MEDIUMINT NOT NULL AUTO_INCREMENT,
`name` varchar(100) NOT NULL,
`genre` varchar(100) NOT NULL,
`year` int(10) NOT NULL,
`award` int(100) unsigned NOT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table movie;