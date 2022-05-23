package candles

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

func NewFillingHistoricalCandlesData(
	instrumentSettingRepository repository.InstrumentSettingRepositoryInterface,
	candleRepository repository.CandleRepositoryInterface,
	marketDataService tinkoffinvest.MarketDataServiceInterface,
	logger log.LoggerInterface,
) FillingHistoricalCandlesDataInterface {
	return &FillingHistoricalCandlesData{
		instrumentSettingRepository: instrumentSettingRepository,
		candleRepository:            candleRepository,
		marketDataService:           marketDataService,
		logger:                      logger,
	}
}

type FillingHistoricalCandlesData struct {
	instrumentSettingRepository repository.InstrumentSettingRepositoryInterface
	candleRepository            repository.CandleRepositoryInterface
	marketDataService           tinkoffinvest.MarketDataServiceInterface
	logger                      log.LoggerInterface
}

func (fhcd FillingHistoricalCandlesData) FillingHistoricalData(ctx context.Context) {
	collectingInstruments, err := fhcd.instrumentSettingRepository.GetInstrumentsSettingsWhereIsCollectingTrue(ctx)

}
