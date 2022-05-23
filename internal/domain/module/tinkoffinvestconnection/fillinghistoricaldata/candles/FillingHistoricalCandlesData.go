package candles

import (
	"context"
	"fmt"
	"time"

	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/dateoperation"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc/investapi"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
	"github.com/shopspring/decimal"
)

func NewFillingHistoricalCandlesData(
	instrumentSettingRepository repository.InstrumentSettingRepositoryInterface,
	candleRepository repository.CandleRepositoryInterface,
	marketDataService tinkoffinvest.MarketDataServiceInterface,
	dateOperation dateoperation.DateOperationInterface,
	config *configs.Config,
	logger log.LoggerInterface,
) FillingHistoricalCandlesDataInterface {
	return &FillingHistoricalCandlesData{
		instrumentSettingRepository: instrumentSettingRepository,
		candleRepository:            candleRepository,
		marketDataService:           marketDataService,
		dateOperation:               dateOperation,
		config:                      config,
		logger:                      logger,
	}
}

type FillingHistoricalCandlesData struct {
	instrumentSettingRepository repository.InstrumentSettingRepositoryInterface
	candleRepository            repository.CandleRepositoryInterface
	marketDataService           tinkoffinvest.MarketDataServiceInterface
	dateOperation               dateoperation.DateOperationInterface
	config                      *configs.Config
	logger                      log.LoggerInterface
}

func (fhcd FillingHistoricalCandlesData) FillingHistoricalData(ctx context.Context) {
	collectingInstruments, err := fhcd.instrumentSettingRepository.GetInstrumentsSettingsWhereIsCollectingTrue(ctx)
	if err != nil {
		return
	}

	for i := 0; i < len(collectingInstruments); i++ {
		fhcd.fillingData(ctx, collectingInstruments[i].Figi)
	}
}

func (fhcd FillingHistoricalCandlesData) FillingHistoricalDataByFigi(ctx context.Context, figi string) {
	instrumentSetting, err := fhcd.instrumentSettingRepository.GetInstrumentSettingByFigi(ctx, figi)
	if err != nil {
		return
	}

	if instrumentSetting.IsDataCollecting == false {
		fhcd.logger.Warning(log.LogCategoryLogic, "Instrument is not data collecting", map[string]interface{}{"service": "FillingHistoricalCandlesData", "method": "FillingHistoricalDataByFigi", "action": "CheckIsDataCollecting"})
	}

	fhcd.fillingData(ctx, figi)
}

func (fhcd FillingHistoricalCandlesData) fillingData(ctx context.Context, figi string) {
	lastCandle, err := fhcd.candleRepository.GetLastCandleByFigi(ctx, figi)
	if err != nil {
		return
	}
	// для получения исторических данных, мы будем двигаться
	// от настоящего к прошлому, пока значение не дойдет до
	// отметки обозначенного в конфиге PeriodMonthForGetHistoricalData
	// или до значения последней свечи, которая уже есть у нас в базе
	timeTo := time.Now()
	monthPeriod := fhcd.config.ApplicationConfig.PeriodMonthForGetHistoricalData
	maxTimeFrom := time.Now().AddDate(0, -monthPeriod, 0)

	for timeTo.After(maxTimeFrom) && timeTo.After(lastCandle.Timestamp) {
		timeFrom := timeTo.AddDate(0, 0, -1)
		candles, err := fhcd.marketDataService.GetHistoricalCandlesByFigi(ctx, figi, timeFrom, timeTo)
		if err != nil {
			fhcd.logger.Error(log.LogCategoryGrpcTinkoff, err.Error(), map[string]interface{}{"service": "FillingHistoricalCandlesData", "method": "fillingData", "action": "getRotaionsHistoricalData"})
			continue
		}
		fhcd.saveCandles(ctx, figi, candles.Candles)
		timeTo = timeFrom
	}
}

func (fhcd FillingHistoricalCandlesData) saveCandles(ctx context.Context, figi string, candles []*investapi.HistoricCandle) {
	for i := 0; i < len(candles); i++ {
		open, openErr := decimal.NewFromString(fmt.Sprintf("%d.%d", candles[i].Open.Units, candles[i].Open.Nano))
		high, highErr := decimal.NewFromString(fmt.Sprintf("%d.%d", candles[i].High.Units, candles[i].High.Nano))
		low, lowErr := decimal.NewFromString(fmt.Sprintf("%d.%d", candles[i].Low.Units, candles[i].Low.Nano))
		close, closeErr := decimal.NewFromString(fmt.Sprintf("%d.%d", candles[i].Close.Units, candles[i].Close.Nano))

		if openErr != nil || highErr != nil || lowErr != nil || closeErr != nil {
			fhcd.logger.Error(log.LogCategoryGrpcTinkoff, "Problem with handle historical candle", map[string]interface{}{"service": "FillingHistoricalCandlesData", "method": "saveCandles", "action": "saveCandles"})
		}
		candle := model.NewCandle(figi, open, high, low, close, candles[i].Volume, candles[i].Time.AsTime())

		fhcd.candleRepository.Save(ctx, candle)
	}
}
