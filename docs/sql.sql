CREATE TABLE `member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `nickname` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `gender` char(16) NOT NULL,
  `mobile` varchar(128) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `status` char(32) DEFAULT NULL,
  `last_ip` varchar(128) DEFAULT NULL,
  `last_login` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 

CREATE TABLE `friend` (
    `id` int(11) not null auto_increment,
    `name` varchar(64) not null,
    `friend` varchar(64) not null,
    `memo` varchar(64) not null,
    `ban` boolean default false,
    `last_update` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`)
);


CREATE TABLE `add_request` (
    `id` int(11) not null auto_increment,
    `owner` varchar(64) not null,
    `requester` varchar(64) not null,
    `message` varchar(64),
    `status` varchar(64) not null,
    `last_update` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`)
);