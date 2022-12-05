CREATE TABLE `comments` (
    `id` bigint PRIMARY KEY AUTO_INCREMENT,
    `article_id` bigint NOT NULL,
    `parent_id` bigint NOT NULL DEFAULT 0,
    `username` varchar(255) NOT NULL,
    `content` text NOT NULL,
    `created_at` datetime,
    `updated_at` datetime,
    `deleted_at` datetime
);

ALTER TABLE `comments` ADD FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`);
