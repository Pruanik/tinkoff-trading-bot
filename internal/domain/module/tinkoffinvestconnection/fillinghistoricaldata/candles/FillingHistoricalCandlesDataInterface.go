package candles

import "context"

type FillingHistoricalCandlesDataInterface interface {
	FillingHistoricalData(ctx context.Context)

	FillingHistoricalDataByFigi(ctx context.Context, figi string)
}
