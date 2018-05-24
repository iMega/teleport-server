CREATE DATABASE IF NOT EXISTS `stock`
  COLLATE 'utf8mb4_unicode_ci'
  DEFAULT CHARSET 'utf8mb4';

USE stock;

CREATE TABLE `entities` (
  `owner_id`    BINARY(16)   NOT NULL COMMENT 'Идентификатор владельца',
  `entity_id`   BINARY(16)   NOT NULL COMMENT 'Идентификатор сущности',
  `entity_type` VARCHAR(255) NOT NULL COMMENT 'Тип сущности',
  `entity`      JSON         NOT NULL COMMENT 'сущность',
  `created`     DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Дата создания записи',
  `updated`     DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Дата обновления записи',
  `deleted`     TINYINT(2)   NOT NULL DEFAULT 0 COMMENT 'состояние удален',
  PRIMARY KEY `id` (`owner_id`, `entity_id`),
  UNIQUE KEY `entity_id` (`entity_id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT 'Сущности';

CREATE TABLE `relations` (
  `owner_id`   BINARY(16)          NOT NULL COMMENT 'Идентификатор сайта',
  `subject_id` BINARY(16)          NOT NULL COMMENT 'Идентификатор субъекта',
  `predicate`  VARCHAR(64)         NOT NULL COMMENT 'предикат',
  `object_id`  BINARY(16)          NOT NULL COMMENT 'Идентификатор объекта',
  `priority`   BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Порядок выдачи по умолчанию',
  `deleted`    TINYINT(2)          NOT NULL DEFAULT 0,
  PRIMARY KEY `id` (`owner_id`, `object_id`,`predicate`, `deleted`),
  KEY priority_index (`owner_id`, `subject_id`,`predicate`,`deleted`, `priority`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COMMENT 'Связи между сущностями';
