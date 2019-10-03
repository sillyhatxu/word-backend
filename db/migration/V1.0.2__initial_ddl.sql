CREATE TABLE IF NOT EXISTS `longman_word`
(
  `id`                 bigint(48)   NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `vocabulary_id`      bigint(48)   NOT NULL,
  `description`        TEXT         NULL,
  `created_time`       timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  `last_modified_time` timestamp(3) NOT NULL DEFAULT current_timestamp(3) ON UPDATE current_timestamp(3),
  INDEX (`vocabulary_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;