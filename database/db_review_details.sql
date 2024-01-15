-- Table: reviewer_details

CREATE TABLE `reviewer_details` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `unix_id` CHAR(12),
  `user_reviewer_id` CHAR(12),
  `campaign_detail_id` CHAR(12),
  `name_reviewer` VARCHAR(255),
  `description` TEXT,
  `status_review` VARCHAR(10),
  `rating` INT(5),
  `updateId_admin` CHAR(12),
  `updateAt_admin` DATETIME,
  `created_at` DATETIME,
  `updated_at` DATETIME,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
