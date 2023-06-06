CREATE TABLE IF NOT EXISTS users
(
    `id`        VARCHAR(255)    NOT NULL  PRIMARY KEY,
    `user_name`  VARCHAR(255) NOT NULL,
    `password`   VARCHAR(255) NOT NULL,
    `email`      VARCHAR(255) NOT NULL UNIQUE,
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP    NULL     DEFAULT NULL
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS wallets
(
    `id`        VARCHAR(255)  NOT NULL  PRIMARY KEY,
    `name`      VARCHAR(255)  NOT NULL,
    `user_id`   VARCHAR(255)  NOT NULL,
    `balance`   integer  NOT NULL DEFAULT 0,
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP    NULL     DEFAULT NULL
);

CREATE TABLE if not exists `refresh_token`
(
    `id`          varchar(64) PRIMARY KEY  COMMENT 'uuid' ,
    `token`       text,
    `user_id`     varchar(255),
    `is_expired`  tinyint(1) DEFAULT 0
    );

CREATE INDEX `idx_user_id` ON `refresh_token` (`user_id`);

ALTER TABLE `refresh_token` ADD CONSTRAINT uq_token UNIQUE(`user_id`);

ALTER TABLE `wallets`
    ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);