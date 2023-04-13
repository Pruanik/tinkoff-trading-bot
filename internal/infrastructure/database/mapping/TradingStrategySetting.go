package mapping

type TradingStrategySetting struct {
}

func (tss *TradingStrategySetting) TableName() string {
	return "trading_strategy_settings"
}
