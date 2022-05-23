package tinkoffinvest

import (
	"context"
	"time"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc/investapi"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewMarketDataService(client investapi.MarketDataServiceClient, logger log.LoggerInterface) tinkoffinvest.MarketDataServiceInterface {
	return &MarketDataService{client: client, logger: logger}
}

type MarketDataService struct {
	client investapi.MarketDataServiceClient
	logger log.LoggerInterface
}

func (mds MarketDataService) GetHistoricalCandlesByFigi(ctx context.Context, figi string, from time.Time, to time.Time) (*investapi.GetCandlesResponse, error) {
	getCandlesRequest := investapi.GetCandlesRequest{
		Figi:     figi,
		From:     timestamppb.New(from),
		To:       timestamppb.New(to),
		Interval: investapi.CandleInterval_CANDLE_INTERVAL_HOUR,
	}
	candles, err := mds.client.GetCandles(ctx, &getCandlesRequest)
	if err != nil {
		mds.logger.Error(log.LogCategoryGrpcTinkoff, err.Error(), map[string]interface{}{"service": "MarketDataService", "method": "GetHistoricalCandlesByFigi"})
		return nil, err
	}

	return candles, nil
}
