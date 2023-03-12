CREATE SCHEMA `enum`;

CREATE TABLE `users` (
  `id` uint PRIMARY KEY,
  `first_name` varchar(50),
  `last_name` varchar(50),
  `phone_number` varchar(15) UNIQUE,
  `avatar` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `personal_messages` (
  `id` uint PRIMARY KEY,
  `sender_id` uint,
  `reciever_id` uint,
  `text` varchar(255),
  `file` varchar(255),
  `status` ENUM ('read', 'sent', 'pending'),
  `sent_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `groups` (
  `id` uint PRIMARY KEY,
  `admin_id` uint,
  `group_name` varchar(50),
  `created_at` timestamp
);

CREATE TABLE `group_members` (
  `id` uint PRIMARY KEY,
  `group_id` uint,
  `user_id` uint,
  `role` ENUM ('member', 'admin'),
  `join_at` timestamp
);

CREATE TABLE `group_messages` (
  `id` uint PRIMARY KEY,
  `sender_member_id` uint,
  `group_id` uint,
  `text` varchar(255),
  `file` varchar(255),
  `status` ENUM ('read', 'sent', 'pending'),
  `sent_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

ALTER TABLE `personal_messages` ADD FOREIGN KEY (`sender_id`) REFERENCES `users` (`id`);

ALTER TABLE `personal_messages` ADD FOREIGN KEY (`reciever_id`) REFERENCES `users` (`id`);

ALTER TABLE `groups` ADD FOREIGN KEY (`admin_id`) REFERENCES `users` (`id`);

ALTER TABLE `group_members` ADD FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`);

ALTER TABLE `group_members` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `group_messages` ADD FOREIGN KEY (`sender_member_id`) REFERENCES `group_members` (`id`);

ALTER TABLE `group_messages` ADD FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`);
