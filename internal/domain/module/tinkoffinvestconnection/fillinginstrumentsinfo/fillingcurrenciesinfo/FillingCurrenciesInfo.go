package fillingcurrenciesinfo

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

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

func (fci FillingCurrenciesInfo) CheckExistAndLoadInfo(ctx context.Context) {
	areCurrenciesExist, err := fci.instrumentRepository.AreInstrumentsExistByType(ctx, "currency")
	if err != nil {
		return
	}

	if !areCurrenciesExist {
		currencies, err := fci.instrumentService.GetBaseCurrencies(ctx)
		if err != nil {
			return
		}

		instruments := currencies.Instruments

		for i := 1; i < len(instruments); i++ {
			instrument := model.NewInstrument(
				instruments[i].GetFigi(),
				instruments[i].GetName(),
				"currency",
			)
			_, err = fci.instrumentRepository.Save(ctx, instrument)
			if err != nil {
				fci.logger.Error(log.LogCategoryLogic, err.Error(), map[string]interface{}{"service": "FillingCurrenciesInfo", "method": "CheckExistAndLoadInfo", "action": "save instrument"})
			}

			currency := model.NewCurrency(
				instruments[i].GetFigi(),
				instruments[i].GetTicker(),
				instruments[i].GetClassCode(),
				instruments[i].GetIsin(),
				instruments[i].GetLot(),
				instruments[i].GetCurrency(),
				instruments[i].GetName(),
				instruments[i].GetExchange(),
				instruments[i].GetOtcFlag(),
				instruments[i].GetBuyAvailableFlag(),
				instruments[i].GetSellAvailableFlag(),
				instruments[i].GetIsoCurrencyName(),
				instruments[i].GetMinPriceIncrement().GetUnits(),
				instruments[i].GetMinPriceIncrement().GetNano(),
				instruments[i].GetApiTradeAvailableFlag(),
			)

			_, err = fci.currenciesRepository.Save(ctx, currency)
			if err != nil {
				fci.logger.Error(log.LogCategoryLogic, err.Error(), map[string]interface{}{"service": "FillingCurrenciesInfo", "method": "CheckExistAndLoadInfo", "action": "save currency"})
			}
		}
	} else {
		fci.logger.Info(log.LogCategoryLogic, "Service does not need filling currencies", make(map[string]interface{}))
	}
}
