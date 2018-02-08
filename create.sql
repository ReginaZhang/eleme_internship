
CREATE TABLE `inventory` (
	`id`		BIGINT(20)	NOT NULL AUTO_INCREMENT             	COMMENT '主键',
                                  	PRIMARY KEY (`id`),
	`created_at`	TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '创建时间',
				  	KEY ix_created_at (`created_at`),
	`updated_at`	TIMESTAMP 	NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '更新时间'
					ON UPDATE CURRENT_TIMESTAMP ,
					KEY ix_updated_at (`updated_at`),
	`drc_check_time` TIMESTAMP(3)   NOT NULL DEFAULT CURRENT_TIMESTAMP(3) 	COMMENT '仅供DRC数据校验使用'
					ON UPDATE CURRENT_TIMESTAMP(3),
					KEY ix_drc_check_time (`drc_check_time`),
	`name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'inventory的显示名称',
	`env`         	VARCHAR(255)	NOT NULL DEFAULT ''                 	COMMENT '环境',
	`version`	VARCHAR(255)	NOT NULL DEFAULT ''			COMMENT '版本号',
	`hosts`		MEDIUMTEXT			NOT NULL			COMMENT '',
	`vars`		MEDIUMTEXT			NOT NULL			COMMENT ''
);

CREATE TABLE `job` (
	`id`		BIGINT(20)	NOT NULL AUTO_INCREMENT             	COMMENT '主键',
                                  	PRIMARY KEY (`id`),
	`created_at`	TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '创建时间',
				  	KEY ix_created_at (`created_at`),
	`updated_at`  TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '更新时间'
					ON UPDATE CURRENT_TIMESTAMP ,
					KEY ix_updated_at (`updated_at`),
	`drc_check_time` TIMESTAMP(3)   NOT NULL DEFAULT CURRENT_TIMESTAMP(3) 	COMMENT '仅供DRC数据校验使用'
					ON UPDATE CURRENT_TIMESTAMP(3),
					KEY ix_drc_check_time (`drc_check_time`),
	`uuid`  VARCHAR(255)  NOT NULL UNIQUE,
	`env`         	VARCHAR(255)	NOT NULL DEFAULT ''                 	COMMENT '环境',
	`playbook_id` BIGINT(20)  NOT NULL DEFAULT  0     COMMENT '对应playbook id',
	`inventory_id`	BIGINT(20)	NOT NULL DEFAULT 0			COMMENT '对应inventory id',
	`status` ENUM('running', 'successful', 'failed') NOT NULL DEFAULT 'running'
) COMMENT='job是一次playbook的执行';

CREATE TABLE `run` (
	`id`		BIGINT(20)	NOT NULL AUTO_INCREMENT             	COMMENT '主键',
                                  	PRIMARY KEY (`id`),
	`created_at`	TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '创建时间',
				  	KEY ix_created_at (`created_at`),
	`updated_at`  TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '更新时间'
					ON UPDATE CURRENT_TIMESTAMP ,
					KEY ix_updated_at (`updated_at`),
	`drc_check_time` TIMESTAMP(3)   NOT NULL DEFAULT CURRENT_TIMESTAMP(3) 	COMMENT '仅供DRC数据校验使用'
					ON UPDATE CURRENT_TIMESTAMP(3),
					KEY ix_drc_check_time (`drc_check_time`),
	`uuid`  VARCHAR(255)  NOT NULL UNIQUE,
	`job_id`       VARCHAR(255)	NOT NULL DEFAULT ''                 	COMMENT '对应job的uuid',
	`env`         	VARCHAR(255)	NOT NULL DEFAULT ''                 	COMMENT '环境',
	`hosts`		MEDIUMTEXT			NOT NULL			COMMENT '',
	`vars`		MEDIUMTEXT			NOT NULL			COMMENT '',
	`limit`   MEDIUMTEXT      NOT NULL      COMMENT ''
) COMMENT='run是一个play中被切分的一部分任务';

CREATE TABLE `app` (
	`id`		BIGINT(20)	NOT NULL AUTO_INCREMENT             	COMMENT '主键',
                                  	PRIMARY KEY (`id`),
	`created_at`	TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '创建时间',
				  	KEY ix_created_at (`created_at`),
	`updated_at`  TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '更新时间'
					ON UPDATE CURRENT_TIMESTAMP ,
					KEY ix_updated_at (`updated_at`),
	`drc_check_time` TIMESTAMP(3)   NOT NULL DEFAULT CURRENT_TIMESTAMP(3) 	COMMENT '仅供DRC数据校验使用'
					ON UPDATE CURRENT_TIMESTAMP(3),
					KEY ix_drc_check_time (`drc_check_time`),
	`appid`       	VARCHAR(255)	NOT NULL DEFAULT ''                 	COMMENT 'appid',
					UNIQUE KEY ux_appid (`appid`)
);

CREATE TABLE `app_playbook`(
  `id`		BIGINT(20)	NOT NULL AUTO_INCREMENT             	COMMENT '主键',
                                  	PRIMARY KEY (`id`),
	`created_at`	TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '创建时间',
				  	KEY ix_created_at (`created_at`),
	`updated_at`  TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '更新时间'
					ON UPDATE CURRENT_TIMESTAMP ,
					KEY ix_updated_at (`updated_at`),
	`drc_check_time` TIMESTAMP(3)   NOT NULL DEFAULT CURRENT_TIMESTAMP(3) 	COMMENT '仅供DRC数据校验使用'
					ON UPDATE CURRENT_TIMESTAMP(3),
					KEY ix_drc_check_time (`drc_check_time`),
	`app_id` BIGINT(20)	NOT NULL DEFAULT  0,
	`playbook_id` BIGINT(20)	NOT NULL DEFAULT  0
);

CREATE TABLE `app_inventory`(
  `id`		BIGINT(20)	NOT NULL AUTO_INCREMENT             	COMMENT '主键',
                                  	PRIMARY KEY (`id`),
	`created_at`	TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '创建时间',
				  	KEY ix_created_at (`created_at`),
	`updated_at`  TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '更新时间'
					ON UPDATE CURRENT_TIMESTAMP ,
					KEY ix_updated_at (`updated_at`),
	`drc_check_time` TIMESTAMP(3)   NOT NULL DEFAULT CURRENT_TIMESTAMP(3) 	COMMENT '仅供DRC数据校验使用'
					ON UPDATE CURRENT_TIMESTAMP(3),
					KEY ix_drc_check_time (`drc_check_time`),
  `app_id` BIGINT(20)	NOT NULL DEFAULT  0,
	`inventory_id` BIGINT(20)	NOT NULL DEFAULT  0
);

CREATE TABLE `statistics` (
  `id`		BIGINT(20)	NOT NULL AUTO_INCREMENT             	COMMENT '主键',
                                  	PRIMARY KEY (`id`),
	`created_at`	TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '创建时间',
				  	KEY ix_created_at (`created_at`),
	`updated_at`  TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '更新时间'
					ON UPDATE CURRENT_TIMESTAMP ,
					KEY ix_updated_at (`updated_at`),
	`drc_check_time` TIMESTAMP(3)   NOT NULL DEFAULT CURRENT_TIMESTAMP(3) 	COMMENT '仅供DRC数据校验使用'
					ON UPDATE CURRENT_TIMESTAMP(3),
					KEY ix_drc_check_time (`drc_check_time`),
	`job_start_time` TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`run_id` VARCHAR(255)	NOT NULL DEFAULT '',
	`target` VARCHAR(255) NOT NULL DEFAULT '',
	`unreachable` TINYINT(1) NOT NULL DEFAULT 0,
	`failed` INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE `failure` (
  `id`		BIGINT(20)	NOT NULL AUTO_INCREMENT             	COMMENT '主键',
                                  	PRIMARY KEY (`id`),
	`created_at`	TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '创建时间',
				  	KEY ix_created_at (`created_at`),
	`updated_at`  TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP  	COMMENT '更新时间'
					ON UPDATE CURRENT_TIMESTAMP ,
					KEY ix_updated_at (`updated_at`),
	`drc_check_time` TIMESTAMP(3)   NOT NULL DEFAULT CURRENT_TIMESTAMP(3) 	COMMENT '仅供DRC数据校验使用'
					ON UPDATE CURRENT_TIMESTAMP(3),
					KEY ix_drc_check_time (`drc_check_time`),
  `stats_id` BIGINT(20)	NOT NULL DEFAULT 0,
  `result` MEDIUMTEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS `pansible`.`playbook` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `drc_check_time` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '仅供drc数据校验使用',
  `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'playbook的显示名称',
  `git_repo` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'playbook所在git仓库的地址',
  `entry` VARCHAR(255) NOT NULL DEFAULT 'play.yml' COMMENT 'playbook入口文件的相对路径',
  PRIMARY KEY (`id`),
  INDEX `ix_created_at` (`created_at` ASC),
  INDEX `ix_updated_at` (`updated_at` ASC),
  INDEX `ix_drc_check_time` (`drc_check_time` ASC));

CREATE TABLE IF NOT EXISTS `pansible`.`run` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `drc_check_time` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '仅供drc数据校验使用',
  `playbook_id` BIGINT(20) NOT NULL DEFAULT 0 COMMENT '使用的playbook id',
  `commit` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '执行使用的git仓库commit',
  `inventory_id` BIGINT(20) NOT NULL DEFAULT 0 COMMENT 'task使用的inventory id',
  `user` VARCHAR(255) NOT NULL DEFAULT 'UNKNOWN' COMMENT '发起task的用户名',
  PRIMARY KEY (`id`),
  INDEX `ix_created_at` (`created_at` ASC),
  INDEX `ix_updated_at` (`updated_at` ASC),
  INDEX `ix_drc_check_time` (`drc_check_time` ASC));
