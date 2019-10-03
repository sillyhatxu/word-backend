CREATE TABLE IF NOT EXISTS `vocabulary`
(
  `id`                 bigint(48)   NOT NULL AUTO_INCREMENT,
  `word`               varchar(100) NOT NULL,
  `created_time`       timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  `last_modified_time` timestamp(3) NOT NULL DEFAULT current_timestamp(3) ON UPDATE current_timestamp(3),
  PRIMARY KEY (`id`),
  INDEX (`word`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;