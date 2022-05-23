package candles

import (
	"context"
	"fmt"
	"time"

	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/dateoperation"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc/investapi"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	monthPeriod := fhcd.config.ApplicationConfig.PeriodMonthForGetHistoricalData

	var timestampFrom *timestamppb.Timestamp
	if lastCandle.Id == 0 || fhcd.dateOperation.MonthsCountSince(lastCandle.Time) > monthPeriod {
		timeFrom := time.Now().AddDate(0, -monthPeriod, 0)
		timestampFrom = timestamppb.New(timeFrom)
	} else {
		timestampFrom = timestamppb.New(lastCandle.Time)
	}

	fhcd.logger.Info(log.LogCategoryGrpcTinkoff, fmt.Sprintf("Load historical data for %s from %s", figi, timestampFrom.String()), make(map[string]interface{}))
	candles, err := fhcd.marketDataService.GetHistoricalCandlesByFigi(ctx, figi, timestampFrom)
	if err != nil {
		return
	}

	fhcd.saveCandles(ctx, figi, candles.Candles)
}

func (fhcd FillingHistoricalCandlesData) saveCandles(ctx context.Context, figi string, candles []*investapi.HistoricCandle) {
	for i := 0; i < len(candles); i++ {
		fmt.Println(figi + " " + candles[i].High.String())
	}
}
