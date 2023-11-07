-- Create "actors" table
CREATE TABLE `actors` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT);
-- Create "albums" table
CREATE TABLE `albums` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `title` text NOT NULL, `year` integer NULL, `description` text NULL, `cover_albums` integer NOT NULL, CONSTRAINT `albums_covers_albums` FOREIGN KEY (`cover_albums`) REFERENCES `covers` (`id`) ON DELETE NO ACTION);
-- Create "artists" table
CREATE TABLE `artists` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL);
-- Create index "artists_name_key" to table: "artists"
CREATE UNIQUE INDEX `artists_name_key` ON `artists` (`name`);
-- Create "audiobooks" table
CREATE TABLE `audiobooks` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT);
-- Create "categories" table
CREATE TABLE `categories` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT);
-- Create "covers" table
CREATE TABLE `covers` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `width` integer NOT NULL, `height` integer NOT NULL, `file_covers` integer NOT NULL, CONSTRAINT `covers_files_covers` FOREIGN KEY (`file_covers`) REFERENCES `files` (`id`) ON DELETE NO ACTION);
-- Create "devices" table
CREATE TABLE `devices` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `ip` text NOT NULL, `device` text NOT NULL, `user_devices` integer NULL, CONSTRAINT `devices_users_devices` FOREIGN KEY (`user_devices`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create "dirs" table
CREATE TABLE `dirs` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `name` text NOT NULL, `user_dirs` integer NULL, CONSTRAINT `dirs_users_dirs` FOREIGN KEY (`user_dirs`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create "directors" table
CREATE TABLE `directors` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT);
-- Create "files" table
CREATE TABLE `files` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `cid` text NOT NULL, `name` text NOT NULL, `size` integer NOT NULL);
-- Create index "files_cid_key" to table: "files"
CREATE UNIQUE INDEX `files_cid_key` ON `files` (`cid`);
-- Create index "file_cid" to table: "files"
CREATE UNIQUE INDEX `file_cid` ON `files` (`cid`);
-- Create "musics" table
CREATE TABLE `musics` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `name` text NOT NULL, `duration` real NOT NULL, `description` text NULL, `file_musics` integer NOT NULL, CONSTRAINT `musics_files_musics` FOREIGN KEY (`file_musics`) REFERENCES `files` (`id`) ON DELETE NO ACTION);
-- Create "playlists" table
CREATE TABLE `playlists` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `name` text NOT NULL, `description` text NULL, `private` bool NOT NULL DEFAULT false, `cover_playlists` integer NULL, `user_playlists` integer NOT NULL, CONSTRAINT `playlists_covers_playlists` FOREIGN KEY (`cover_playlists`) REFERENCES `covers` (`id`) ON DELETE SET NULL, CONSTRAINT `playlists_users_playlists` FOREIGN KEY (`user_playlists`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create "saveds" table
CREATE TABLE `saveds` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `name` text NOT NULL, `caption` text NULL, `file_saves` integer NULL, `user_saves` integer NULL, CONSTRAINT `saveds_files_saves` FOREIGN KEY (`file_saves`) REFERENCES `files` (`id`) ON DELETE SET NULL, CONSTRAINT `saveds_users_saves` FOREIGN KEY (`user_saves`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create "users" table
CREATE TABLE `users` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `username` text NOT NULL, `password` text NULL, `email` text NULL, `name` text NULL, `bio` text NULL, `avatar` text NULL);
-- Create index "users_username_key" to table: "users"
CREATE UNIQUE INDEX `users_username_key` ON `users` (`username`);
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX `users_email_key` ON `users` (`email`);
-- Create "vodeos" table
CREATE TABLE `vodeos` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT);
-- Create "album_musics" table
CREATE TABLE `album_musics` (`album_id` integer NOT NULL, `music_id` integer NOT NULL, PRIMARY KEY (`album_id`, `music_id`), CONSTRAINT `album_musics_album_id` FOREIGN KEY (`album_id`) REFERENCES `albums` (`id`) ON DELETE CASCADE, CONSTRAINT `album_musics_music_id` FOREIGN KEY (`music_id`) REFERENCES `musics` (`id`) ON DELETE CASCADE);
-- Create "artist_musics" table
CREATE TABLE `artist_musics` (`artist_id` integer NOT NULL, `music_id` integer NOT NULL, PRIMARY KEY (`artist_id`, `music_id`), CONSTRAINT `artist_musics_artist_id` FOREIGN KEY (`artist_id`) REFERENCES `artists` (`id`) ON DELETE CASCADE, CONSTRAINT `artist_musics_music_id` FOREIGN KEY (`music_id`) REFERENCES `musics` (`id`) ON DELETE CASCADE);
-- Create "artist_albums" table
CREATE TABLE `artist_albums` (`artist_id` integer NOT NULL, `album_id` integer NOT NULL, PRIMARY KEY (`artist_id`, `album_id`), CONSTRAINT `artist_albums_artist_id` FOREIGN KEY (`artist_id`) REFERENCES `artists` (`id`) ON DELETE CASCADE, CONSTRAINT `artist_albums_album_id` FOREIGN KEY (`album_id`) REFERENCES `albums` (`id`) ON DELETE CASCADE);
-- Create "dir_saves" table
CREATE TABLE `dir_saves` (`dir_id` integer NOT NULL, `saved_id` integer NOT NULL, PRIMARY KEY (`dir_id`, `saved_id`), CONSTRAINT `dir_saves_dir_id` FOREIGN KEY (`dir_id`) REFERENCES `dirs` (`id`) ON DELETE CASCADE, CONSTRAINT `dir_saves_saved_id` FOREIGN KEY (`saved_id`) REFERENCES `saveds` (`id`) ON DELETE CASCADE);
-- Create "dir_subdir" table
CREATE TABLE `dir_subdir` (`dir_id` integer NOT NULL, `pdir_id` integer NOT NULL, PRIMARY KEY (`dir_id`, `pdir_id`), CONSTRAINT `dir_subdir_dir_id` FOREIGN KEY (`dir_id`) REFERENCES `dirs` (`id`) ON DELETE CASCADE, CONSTRAINT `dir_subdir_pdir_id` FOREIGN KEY (`pdir_id`) REFERENCES `dirs` (`id`) ON DELETE CASCADE);
-- Create "playlist_musics" table
CREATE TABLE `playlist_musics` (`playlist_id` integer NOT NULL, `music_id` integer NOT NULL, PRIMARY KEY (`playlist_id`, `music_id`), CONSTRAINT `playlist_musics_playlist_id` FOREIGN KEY (`playlist_id`) REFERENCES `playlists` (`id`) ON DELETE CASCADE, CONSTRAINT `playlist_musics_music_id` FOREIGN KEY (`music_id`) REFERENCES `musics` (`id`) ON DELETE CASCADE);
-- Create "user_albums" table
CREATE TABLE `user_albums` (`user_id` integer NOT NULL, `album_id` integer NOT NULL, PRIMARY KEY (`user_id`, `album_id`), CONSTRAINT `user_albums_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `user_albums_album_id` FOREIGN KEY (`album_id`) REFERENCES `albums` (`id`) ON DELETE CASCADE);
-- Create "user_musics" table
CREATE TABLE `user_musics` (`user_id` integer NOT NULL, `music_id` integer NOT NULL, PRIMARY KEY (`user_id`, `music_id`), CONSTRAINT `user_musics_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `user_musics_music_id` FOREIGN KEY (`music_id`) REFERENCES `musics` (`id`) ON DELETE CASCADE);
