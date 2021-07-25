CREATE DATABASE
    IF
        NOT EXISTS blog_service DEFAULT CHARACTER
        SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

CREATE TABLE `blog_article` (
                                `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
                                `title` VARCHAR(100) DEFAULT '' COMMENT '文章標題',
                                `desc` VARCHAR(255) DEFAULT '' COMMENT '文章簡述',
                                `cover_image_url` VARCHAR(255) DEFAULT '' COMMENT '封面圖片位址',
                                `content` longtext COMMENT '文章內容',
                                `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '建立時間',
                                `created_by` VARCHAR(100) DEFAULT '' COMMENT '建立人',
                                `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改時間',
                                `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
                                `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '刪除時間',
                                `is_del` tinyint(3) DEFAULT '0' COMMENT '是否刪除0為未刪除、1為已刪除',
                                `state` tinyint(3) UNSIGNED DEFAULT '1' COMMENT '狀態0為禁用、1為啟用',
                                PRIMARY KEY (`id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '文章管理';

CREATE TABLE `blog_article_tag` (
                                    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
                                    `article_id` INT(11) NOT NULL COMMENT '文章ID',
                                    `tag_id` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '標籤ID',
                                    `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '建立時間',
                                    `created_by` VARCHAR(100) DEFAULT '' COMMENT '建立人',
                                    `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改時間',
                                    `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
                                    `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '刪除時間',
                                    `is_del` tinyint(3) DEFAULT '0' COMMENT '是否刪除0為未刪除、1為已刪除',
                                    PRIMARY KEY (`id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '文章標籤連結';

CREATE TABLE `blog_auth` (
                             `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
                             `app_key` VARCHAR(20) DEFAULT '' COMMENT 'Key',
                             `app_secret` VARCHAR(50) DEFAULT '' COMMENT 'Serect',
                             `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '建立時間',
                             `created_by` VARCHAR(100) DEFAULT '' COMMENT '建立人',
                             `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改時間',
                             `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
                             `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '刪除時間',
                             `is_del` tinyint(3) DEFAULT '0' COMMENT '是否刪除0為未刪除、1為已刪除',
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '認證管理';

CREATE TABLE `blog_tag` (
                            `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
                            `name` VARCHAR(100) DEFAULT '' COMMENT '標籤名稱',
                            `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '建立時間',
                            `created_by` VARCHAR(100) DEFAULT '' COMMENT '建立人',
                            `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改時間',
                            `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
                            `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '刪除時間',
                            `is_del` tinyint(3) DEFAULT '0' COMMENT '是否刪除0為未刪除、1為已刪除',
                            `state` tinyint(3) UNSIGNED DEFAULT '1' COMMENT '狀態0為禁用、1為啟用',
                            PRIMARY KEY (`id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '標籤管理';

CREATE TABLE `blog_tag` (
                            `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
                            `name` VARCHAR(100) DEFAULT '' COMMENT '標籤名稱',
                            `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '建立時間',
                            `created_by` VARCHAR(100) DEFAULT '' COMMENT '建立人',
                            `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改時間',
                            `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
                            `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '刪除時間',
                            `is_del` tinyint(3) DEFAULT '0' COMMENT '是否刪除0為未刪除、1為已刪除',
                            `state` tinyint(3) UNSIGNED DEFAULT '1' COMMENT '狀態0為禁用、1為啟用',
                            PRIMARY KEY (`id`)
) ENGINE = INNODB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8mb4 COMMENT = '標籤管理';

CREATE TABLE `blog_auth` (
                             `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
                             `app_key` VARCHAR(20) DEFAULT '' COMMENT 'Key',
                             `app_secret` VARCHAR(50) DEFAULT '' COMMENT 'Serect',
                             `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '建立時間',
                             `created_by` VARCHAR(100) DEFAULT '' COMMENT '建立人',
                             `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改時間',
                             `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
                             `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '刪除時間',
                             `is_del` tinyint(3) DEFAULT '0' COMMENT '是否刪除0為未刪除、1為已刪除',
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '認證管理';
