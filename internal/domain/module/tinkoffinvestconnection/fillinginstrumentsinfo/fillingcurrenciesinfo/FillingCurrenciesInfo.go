package fillingcurrenciesinfo

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc/investapi"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

const currencySectorName string = "currency"
const instrumentType string = "currency"

func NewFillingCurrenciesInfo(
	instrumentRepository repository.InstrumentRepositoryInterface,
	currenciesRepository repository.CurrencyRepositoryInterface,
	instrumentService tinkoffinvest.InstrumentServiceInterface,
	logger log.LoggerInterface,
) FillingCurrenciesInfoInterface {
	return &FillingCurrenciesInfo{
		instrumentRepository: instrumentRepository,
		currenciesRepository: currenciesRepository,
		instrumentService:    instrumentService,
		logger:               logger,
	}
}

type FillingCurrenciesInfo struct {
	instrumentRepository repository.InstrumentRepositoryInterface
	currenciesRepository repository.CurrencyRepositoryInterface
	instrumentService    tinkoffinvest.InstrumentServiceInterface
	logger               log.LoggerInterface
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

		fci.createInstrument(ctx, instruments[i])
	}
}

func (fci FillingCurrenciesInfo) createInstrument(ctx context.Context, instrumentCurrency *investapi.Currency) {
	instrument := model.NewInstrument(
		instrumentCurrency.GetFigi(),
		instrumentCurrency.GetName(),
		currencySectorName,
		instrumentType,
	)

	_, err := fci.instrumentRepository.Save(ctx, instrument)
	if err != nil {
		fci.logger.Error(
			log.LogCategoryLogic,
			err.Error(),
			map[string]interface{}{"service": "FillingCurrenciesInfo", "method": "CreateInstrumentsIfNotExist", "action": "save instrument"},
		)
	}

	currency := model.NewCurrency(
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

	_, err = fci.currenciesRepository.Save(ctx, currency)
	if err != nil {
		fci.logger.Error(
			log.LogCategoryLogic,
			err.Error(),
			map[string]interface{}{"service": "FillingCurrenciesInfo", "method": "CreateInstrumentsIfNotExist", "action": "save currency"},
		)
	}
}
