CREATE TABLE `articles` (
    `id` bigint PRIMARY KEY AUTO_INCREMENT,
    `username` varchar(255) NOT NULL,
    `title` varchar(255) NOT NULL,
    `slug` varchar(255) NOT NULL,
    `content` text NOT NULL,
    `thumbnail` text,
    `created_at` datetime,
    `updated_at` datetime,
    `deleted_at` datetime
);