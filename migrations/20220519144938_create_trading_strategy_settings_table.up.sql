CREATE TABLE IF NOT EXISTS `trading_strategy_settings` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`trading_strategy_id` INT NOT NULL,
	`notes` TEXT,
	`parameters` jsonb,
	`created_at` DATETIME NOT NULL,
    KEY `i_trading_strategy_settings_trading_strategy_id` (`trading_strategy_id`) USING BTREE,
	PRIMARY KEY (`id`),
	CONSTRAINT fk_trading_strategy_settings_trading_strategy_id
      FOREIGN KEY(trading_strategy_id) 
	  REFERENCES trading_strategies(id)
);
