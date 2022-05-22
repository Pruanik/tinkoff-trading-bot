package fillingcurrenciesinfo

import (
	"context"
	"fmt"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
)

func NewFillingCurrenciesInfo(
	instrumentRepository repository.InstrumentRepositoryInterface,
	currenciesRepository repository.CurrencyRepositoryInterface,
	instrumentService tinkoffinvest.InstrumentServiceInterface,
) FillingCurrenciesInfoInterface {
	return &FillingCurrenciesInfo{
		instrumentRepository: instrumentRepository,
		currenciesRepository: currenciesRepository,
		instrumentService:    instrumentService,
	}
}

type FillingCurrenciesInfo struct {
	instrumentRepository repository.InstrumentRepositoryInterface
	currenciesRepository repository.CurrencyRepositoryInterface
	instrumentService    tinkoffinvest.InstrumentServiceInterface
}

func (fsi FillingCurrenciesInfo) CheckExistAndLoadInfo(ctx context.Context) {
	areCurrenciesExist, err := fsi.instrumentRepository.AreInstrumentsExistByType(ctx, "currency")
	if err != nil {
		fmt.Println("checkDataExistAndLoad Error: " + err.Error())
	}

	if !areCurrenciesExist {
		currencies, err := fsi.instrumentService.GetBaseCurrencies(ctx)
		if err != nil {
			fmt.Println("GetBaseCurrencies error")
		}

		instruments := currencies.Instruments

		for i := 1; i < len(instruments); i++ {
			instrument := model.NewInstrument(
				instruments[i].GetFigi(),
				instruments[i].GetName(),
				"currency",
			)
			_, err = fsi.instrumentRepository.Save(ctx, instrument)
			if err != nil {
				fmt.Println("instrumentRepository.Save error")
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

			_, err = fsi.currenciesRepository.Save(ctx, currency)
			if err != nil {
				fmt.Println("currenciesRepository.Save error")
			}
		}
	}
}
