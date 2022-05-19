CREATE TABLE IF NOT EXISTS `trading_strategies` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR NOT NULL,
	`description` TEXT,
	`created_at` DATETIME NOT NULL,
    KEY `i_candles_created_at` (`created_at`) USING BTREE,
	PRIMARY KEY (`id`)
);
