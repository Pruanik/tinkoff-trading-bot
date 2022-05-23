package fillingsharesinfo

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

func NewFillingSharesInfo(
	instrumentRepository repository.InstrumentRepositoryInterface,
	shareRepository repository.ShareRepositoryInterface,
	instrumentService tinkoffinvest.InstrumentServiceInterface,
	logger log.LoggerInterface,
) FillingSharesInfoInterface {
	return &FillingSharesInfo{
		instrumentRepository: instrumentRepository,
		shareRepository:      shareRepository,
		instrumentService:    instrumentService,
		logger:               logger,
	}
}

type FillingSharesInfo struct {
	instrumentRepository repository.InstrumentRepositoryInterface
	shareRepository      repository.ShareRepositoryInterface
	instrumentService    tinkoffinvest.InstrumentServiceInterface
	logger               log.LoggerInterface
}

func (fsi FillingSharesInfo) CheckExistAndLoadInfo(ctx context.Context) {
	areSharesExist, err := fsi.instrumentRepository.AreInstrumentsExistByType(ctx, "share")
	if err != nil {
		return
	}

	if !areSharesExist {
		shares, err := fsi.instrumentService.GetBaseShares(ctx)
		if err != nil {
			return
		}

		instruments := shares.Instruments

		for i := 0; i < len(instruments); i++ {
			instrument := model.NewInstrument(
				instruments[i].GetFigi(),
				instruments[i].GetName(),
				"share",
			)
			_, err = fsi.instrumentRepository.Save(ctx, instrument)
			if err != nil {
				fsi.logger.Error(log.LogCategoryLogic, err.Error(), map[string]interface{}{"service": "FillingSharesInfo", "method": "CheckExistAndLoadInfo", "action": "save instrument"})
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
				fsi.logger.Error(log.LogCategoryLogic, err.Error(), map[string]interface{}{"service": "FillingSharesInfo", "method": "CheckExistAndLoadInfo", "action": "save share"})
			}
		}
	} else {
		fsi.logger.Info(log.LogCategoryLogic, "Service does not need filling shares", make(map[string]interface{}))
	}
}
