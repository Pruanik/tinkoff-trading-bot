package mapping

type TradingStrategy struct {
}

func (ts *TradingStrategy) TableName() string {
	return "trading_strategies"
}
