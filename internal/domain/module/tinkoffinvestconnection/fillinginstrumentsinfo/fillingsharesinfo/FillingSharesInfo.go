package fillingsharesinfo

import (
	"context"
	"fmt"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
)

func NewFillingSharesInfo(
	instrumentRepository repository.InstrumentRepositoryInterface,
	shareRepository repository.ShareRepositoryInterface,
	instrumentService tinkoffinvest.InstrumentServiceInterface,
) FillingSharesInfoInterface {
	return &FillingSharesInfo{
		instrumentRepository: instrumentRepository,
		shareRepository:      shareRepository,
		instrumentService:    instrumentService,
	}
}

type FillingSharesInfo struct {
	instrumentRepository repository.InstrumentRepositoryInterface
	shareRepository      repository.ShareRepositoryInterface
	instrumentService    tinkoffinvest.InstrumentServiceInterface
}

func (fsi FillingSharesInfo) CheckExistAndLoadInfo(ctx context.Context) {
	areSharesExist, err := fsi.instrumentRepository.AreInstrumentsExistByType(ctx, "share")
	if err != nil {
		fmt.Println("checkDataExistAndLoad Error: " + err.Error())
	}

	if !areSharesExist {
		shares, err := fsi.instrumentService.GetBaseShares(ctx)
		if err != nil {
			fmt.Println("GetBaseShares error")
		}

		instruments := shares.Instruments

		for i := 1; i < len(instruments); i++ {
			instrument := model.NewInstrument(
				instruments[i].GetFigi(),
				instruments[i].GetName(),
				"share",
			)
			_, err = fsi.instrumentRepository.Save(ctx, instrument)
			if err != nil {
				fmt.Println("instrumentRepository.Save error")
			}

			share := model.NewShare(
				instruments[i].GetFigi(),
				instruments[i].GetTicker(),
				instruments[i].GetClassCode(),
				instruments[i].GetIsin(),
				instruments[i].GetLot(),
				instruments[i].GetCurrency(),
				instruments[i].GetName(),
				instruments[i].GetExchange(),
				instruments[i].GetSector(),
				instruments[i].GetMinPriceIncrement().GetUnits(),
				instruments[i].GetMinPriceIncrement().GetNano(),
				instruments[i].GetApiTradeAvailableFlag(),
			)

			_, err = fsi.shareRepository.Save(ctx, share)
			if err != nil {
				fmt.Println("shareRepository.Save error")
			}
		}
	}
}
