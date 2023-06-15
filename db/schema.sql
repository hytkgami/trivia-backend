CREATE TABLE `users` (
    `uid` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `lobbies` (
    `lobby_id` varchar(26) NOT NULL,
    `owner_uid` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `is_public` tinyint NOT NULL DEFAULT 0,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`lobby_id`),
    FOREIGN KEY (`owner_uid`) REFERENCES `users`(`uid`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `lobby_status` (
    `lobby_status_id` varchar(26) NOT NULL,
    `status` varchar(255) NOT NULL, -- ('waiting', 'active', 'finished')
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`lobby_status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `lobby_lobby_status` (
    `lobby_id` varchar(26) NOT NULL,
    `lobby_status_id` varchar(26) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`lobby_id`, `lobby_status_id`),
    FOREIGN KEY (`lobby_id`) REFERENCES `lobbies`(`lobby_id`) ON UPDATE CASCADE,
    FOREIGN KEY (`lobby_status_id`) REFERENCES `lobby_status`(`lobby_status_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `questions` (
    `question_id` varchar(26) NOT NULL,
    `created_by` varchar(255) NOT NULL, -- uid of user who created the question
    `lobby_id` varchar(26) NOT NULL,
    `title` text(2048) NOT NULL,
    `order_number` int NOT NULL,
    `score` int NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`question_id`),
    FOREIGN KEY (`created_by`) REFERENCES `users`(`uid`) ON UPDATE CASCADE,
    FOREIGN KEY (`lobby_id`) REFERENCES `lobbies`(`lobby_id`) ON UPDATE CASCADE,
    UNIQUE KEY `uk_lobby_id_order_number_on_questions` (`lobby_id`, `order_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `answers` (
    `answer_id` varchar(26) NOT NULL,
    `question_id` varchar(26) NOT NULL,
    `uid` varchar(255) NOT NULL,
    `content` varchar(255) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`answer_id`),
    FOREIGN KEY (`question_id`) REFERENCES `questions`(`question_id`) ON UPDATE CASCADE,
    FOREIGN KEY (`uid`) REFERENCES `users`(`uid`) ON UPDATE CASCADE,
    UNIQUE (`question_id`, `uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `scores` (
    `answer_id` varchar(26) NOT NULL,
    `mark` enum('correct', 'neutral' ,'incorrect') NOT NULL,
    `value` int NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`answer_id`),
    FOREIGN KEY (`answer_id`) REFERENCES `answers`(`answer_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
