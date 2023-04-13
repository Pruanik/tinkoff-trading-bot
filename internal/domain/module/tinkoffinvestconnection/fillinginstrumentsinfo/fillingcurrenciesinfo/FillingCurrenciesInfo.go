package fillingcurrenciesinfo

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/service"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc/investapi"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

const currencySectorName string = "currency"
const instrumentType string = "currency"

func NewFillingCurrenciesInfo(
	instrumentSectorService service.InstrumentSectorServiceInterface,
	instrumentRepository repository.InstrumentRepositoryInterface,
	currenciesRepository repository.CurrencyRepositoryInterface,
	instrumentService tinkoffinvest.InstrumentServiceInterface,
	logger log.LoggerInterface,
) FillingCurrenciesInfoInterface {
	return &FillingCurrenciesInfo{
		instrumentSectorService: instrumentSectorService,
		instrumentRepository:    instrumentRepository,
		currenciesRepository:    currenciesRepository,
		instrumentService:       instrumentService,
		logger:                  logger,
	}
}

type FillingCurrenciesInfo struct {
	instrumentSectorService service.InstrumentSectorServiceInterface
	instrumentRepository    repository.InstrumentRepositoryInterface
	currenciesRepository    repository.CurrencyRepositoryInterface
	instrumentService       tinkoffinvest.InstrumentServiceInterface
	logger                  log.LoggerInterface
}

func (fci FillingCurrenciesInfo) CreateInstrumentsIfNotExist(ctx context.Context) {
	currencies, err := fci.instrumentService.GetBaseCurrencies(ctx)
	if err != nil {
		return
	}

	instruments := currencies.Instruments

	for i := 0; i < len(instruments); i++ {
		dbInstrument, _ := fci.instrumentRepository.GetInstrumentByFigi(ctx, instruments[i].Figi)
		if dbInstrument != nil {
			break
		}

		fci.createInstrumentCurrency(ctx, instruments[i])
	}
}

func (fci FillingCurrenciesInfo) createInstrumentCurrency(ctx context.Context, instrumentCurrency *investapi.Currency) {
	sector, sectorErr := fci.createSector(ctx, currencySectorName)
	if sectorErr != nil {
		fci.logger.Error(
			log.LogCategoryLogic,
			sectorErr.Error(),
			map[string]interface{}{"service": "FillingCurrenciesInfo", "method": "createSector", "action": "save sector"},
		)
		return
	}

	_, instrumentErr := fci.createInstrument(ctx, instrumentCurrency, sector)
	if instrumentErr != nil {
		fci.logger.Error(
			log.LogCategoryLogic,
			instrumentErr.Error(),
			map[string]interface{}{"service": "FillingCurrenciesInfo", "method": "createInstrument", "action": "save instrument"},
		)
	}

	_, shareErr := fci.createCurrency(ctx, instrumentCurrency)
	if shareErr != nil {
		fci.logger.Error(
			log.LogCategoryLogic,
			shareErr.Error(),
			map[string]interface{}{"service": "FillingCurrenciesInfo", "method": "createShare", "action": "save share"},
		)
	}
}

func (fci FillingCurrenciesInfo) createSector(ctx context.Context, code string) (*model.InstrumentSector, error) {
	return fci.instrumentSectorService.CreateIfNotExist(ctx, code)
}

func (fci FillingCurrenciesInfo) createInstrument(
	ctx context.Context,
	instrumentCurrency *investapi.Currency,
	instrumentSector *model.InstrumentSector,
) (*model.Instrument, error) {
	newInstrument := model.NewInstrument(
		instrumentCurrency.GetFigi(),
		instrumentCurrency.GetName(),
		instrumentSector.Id,
		instrumentType,
	)

	return fci.instrumentRepository.Save(ctx, newInstrument)
}

func (fci FillingCurrenciesInfo) createCurrency(ctx context.Context, instrumentCurrency *investapi.Currency) (*model.Currency, error) {
	newCurrency := model.NewCurrency(
		instrumentCurrency.GetFigi(),
		instrumentCurrency.GetTicker(),
		instrumentCurrency.GetClassCode(),
		instrumentCurrency.GetIsin(),
		instrumentCurrency.GetLot(),
		instrumentCurrency.GetCurrency(),
		instrumentCurrency.GetName(),
		instrumentCurrency.GetExchange(),
		instrumentCurrency.GetOtcFlag(),
		instrumentCurrency.GetBuyAvailableFlag(),
		instrumentCurrency.GetSellAvailableFlag(),
		instrumentCurrency.GetIsoCurrencyName(),
		instrumentCurrency.GetMinPriceIncrement().GetUnits(),
		instrumentCurrency.GetMinPriceIncrement().GetNano(),
		instrumentCurrency.GetApiTradeAvailableFlag(),
	)

	return fci.currenciesRepository.Save(ctx, newCurrency)
}
