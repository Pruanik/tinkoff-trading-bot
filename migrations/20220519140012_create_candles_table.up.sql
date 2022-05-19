CREATE TABLE IF NOT EXISTS `candles` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`figi` VARCHAR NOT NULL,
	`interval` INT,
	`open` DOUBLE NOT NULL,
	`high` DOUBLE NOT NULL,
	`low` DOUBLE NOT NULL,
	`close` DOUBLE NOT NULL,
	`volume` BIGINT NOT NULL,
	`timestamp` BIGINT NOT NULL,
	`created_at` DATETIME NOT NULL,
	KEY `i_candles_figi` (`figi`) USING BTREE,
    KEY `i_candles_created_at` (`created_at`) USING BTREE,
    KEY `i_candles_timestamp` (`timestamp`) USING BTREE,
	PRIMARY KEY (`id`),
    CONSTRAINT fk_candles_instruments_figi
      FOREIGN KEY(figi) 
	  REFERENCES instruments(figi)
);
