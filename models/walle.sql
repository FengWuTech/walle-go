# ************************************************************
# Sequel Pro SQL dump
# Version 5446
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 192.168.238.178 (MySQL 5.7.28-log)
# Database: walle
# Generation Time: 2020-08-03 14:12:16 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table environments
# ------------------------------------------------------------

CREATE TABLE `environments` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT 'master',
  `space_id` int(10) NOT NULL DEFAULT '0',
  `status` tinyint(1) DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_spaceId` (`space_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table members
# ------------------------------------------------------------

CREATE TABLE `members` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_id` int(10) DEFAULT '0',
  `source_id` int(10) DEFAULT '0',
  `source_type` varchar(10) DEFAULT '',
  `access_level` varchar(10) DEFAULT '10',
  `status` tinyint(1) DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_source` (`source_type`,`source_id`,`access_level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table menus
# ------------------------------------------------------------

CREATE TABLE `menus` (
  `id` int(15) NOT NULL AUTO_INCREMENT,
  `name_cn` varchar(30) NOT NULL,
  `name_en` varchar(30) NOT NULL,
  `pid` int(6) NOT NULL,
  `type` enum('action','controller','module') DEFAULT 'action',
  `sequence` int(11) DEFAULT '0',
  `role` varchar(10) NOT NULL DEFAULT '',
  `archive` tinyint(1) DEFAULT '0',
  `icon` varchar(30) DEFAULT '',
  `url` varchar(100) DEFAULT '',
  `visible` tinyint(1) DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table projects
# ------------------------------------------------------------

CREATE TABLE `projects` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_id` int(10) NOT NULL,
  `name` varchar(100) DEFAULT 'master',
  `environment_id` int(1) NOT NULL,
  `space_id` int(10) NOT NULL DEFAULT '0',
  `status` tinyint(1) DEFAULT '1',
  `master` varchar(100) NOT NULL DEFAULT '',
  `version` varchar(40) DEFAULT '',
  `excludes` text,
  `target_root` varchar(200) NOT NULL,
  `target_releases` varchar(200) NOT NULL,
  `server_ids` text,
  `task_vars` text,
  `prev_deploy` text,
  `post_deploy` text,
  `prev_release` text,
  `post_release` text,
  `keep_version_num` int(3) NOT NULL DEFAULT '20',
  `repo_url` varchar(200) DEFAULT '',
  `repo_username` varchar(50) DEFAULT '',
  `repo_password` varchar(50) DEFAULT '',
  `repo_mode` varchar(50) DEFAULT 'branch',
  `repo_type` varchar(10) DEFAULT 'git',
  `notice_type` varchar(10) NOT NULL DEFAULT '',
  `notice_hook` text NOT NULL,
  `task_audit` tinyint(1) DEFAULT '0',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_include` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_spaceId` (`space_id`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table records
# ------------------------------------------------------------

CREATE TABLE `records` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `stage` varchar(20) DEFAULT NULL,
  `sequence` int(10) DEFAULT NULL,
  `user_id` int(21) unsigned NOT NULL,
  `task_id` bigint(11) NOT NULL,
  `status` int(3) NOT NULL,
  `host` varchar(200) DEFAULT '',
  `user` varchar(200) DEFAULT '',
  `command` text,
  `success` longtext,
  `error` longtext,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_taskId` (`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table servers
# ------------------------------------------------------------

CREATE TABLE `servers` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '',
  `host` varchar(100) NOT NULL,
  `port` int(1) DEFAULT '22',
  `status` tinyint(1) DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table spaces
# ------------------------------------------------------------

CREATE TABLE `spaces` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_id` int(10) NOT NULL,
  `name` varchar(100) NOT NULL,
  `status` tinyint(1) DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tasks
# ------------------------------------------------------------

CREATE TABLE `tasks` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `user_id` int(21) unsigned NOT NULL,
  `project_id` int(11) NOT NULL,
  `action` int(1) DEFAULT '0',
  `status` tinyint(1) NOT NULL,
  `link_id` varchar(100) DEFAULT '',
  `ex_link_id` varchar(100) DEFAULT '',
  `servers` text,
  `commit_id` varchar(40) DEFAULT '',
  `branch` varchar(100) DEFAULT 'master',
  `tag` varchar(100) DEFAULT '',
  `file_transmission_mode` smallint(3) NOT NULL DEFAULT '1',
  `file_list` longtext,
  `is_rollback` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_name` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_projectId` (`project_id`,`user_id`),
  KEY `idx_userId` (`user_id`,`project_id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table users
# ------------------------------------------------------------

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `is_email_verified` tinyint(1) NOT NULL DEFAULT '0',
  `email` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  `password_hash` varchar(50) DEFAULT NULL,
  `avatar` varchar(100) NOT NULL DEFAULT 'default.jpg',
  `role` varchar(10) NOT NULL DEFAULT '',
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `last_space` int(11) NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`),
  KEY `idx_name` (`username`,`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
