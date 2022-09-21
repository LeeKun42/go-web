CREATE TABLE `user` (
     `id` int unsigned NOT NULL AUTO_INCREMENT,
     `mobile` varchar(20) COLLATE utf8mb4_general_ci NOT NULL,
     `passwd` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
     `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;