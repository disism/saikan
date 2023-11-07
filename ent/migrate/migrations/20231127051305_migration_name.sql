-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_albums" table
CREATE TABLE `new_albums` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `title` text NOT NULL, `year` integer NULL, `description` text NULL, `image_albums` integer NOT NULL, CONSTRAINT `albums_images_albums` FOREIGN KEY (`image_albums`) REFERENCES `images` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "albums" to new temporary table "new_albums"
INSERT INTO `new_albums` (`id`, `create_time`, `update_time`, `title`, `year`, `description`) SELECT `id`, `create_time`, `update_time`, `title`, `year`, `description` FROM `albums`;
-- Drop "albums" table after copying rows
DROP TABLE `albums`;
-- Rename temporary table "new_albums" to "albums"
ALTER TABLE `new_albums` RENAME TO `albums`;
-- Drop "covers" table
DROP TABLE `covers`;
-- Create "new_playlists" table
CREATE TABLE `new_playlists` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `name` text NOT NULL, `description` text NULL, `private` bool NOT NULL DEFAULT false, `image_playlists` integer NULL, `user_playlists` integer NOT NULL, CONSTRAINT `playlists_images_playlists` FOREIGN KEY (`image_playlists`) REFERENCES `images` (`id`) ON DELETE SET NULL, CONSTRAINT `playlists_users_playlists` FOREIGN KEY (`user_playlists`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "playlists" to new temporary table "new_playlists"
INSERT INTO `new_playlists` (`id`, `create_time`, `update_time`, `name`, `description`, `private`, `user_playlists`) SELECT `id`, `create_time`, `update_time`, `name`, `description`, `private`, `user_playlists` FROM `playlists`;
-- Drop "playlists" table after copying rows
DROP TABLE `playlists`;
-- Rename temporary table "new_playlists" to "playlists"
ALTER TABLE `new_playlists` RENAME TO `playlists`;
-- Create "images" table
CREATE TABLE `images` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `width` integer NOT NULL, `height` integer NOT NULL, `file_images` integer NOT NULL, CONSTRAINT `images_files_images` FOREIGN KEY (`file_images`) REFERENCES `files` (`id`) ON DELETE NO ACTION);
-- Create "oidcs" table
CREATE TABLE `oidcs` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `name` text NOT NULL, `configuration_endpoint` text NOT NULL);
-- Create index "oidcs_name_key" to table: "oidcs"
CREATE UNIQUE INDEX `oidcs_name_key` ON `oidcs` (`name`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
