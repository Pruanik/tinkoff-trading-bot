package fillingsharesinfo

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc/investapi"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

const instrumentType string = "share"

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

func (fsi FillingSharesInfo) CreateInstrumentsIfNotExist(ctx context.Context) {
	shares, err := fsi.instrumentService.GetBaseShares(ctx)
	if err != nil {
		return
	}

	instruments := shares.Instruments

	for i := 0; i < len(instruments); i++ {
		dbInstrument, _ := fsi.instrumentRepository.GetInstrumentByFigi(ctx, instruments[i].GetFigi())
		if dbInstrument != nil {
			break
		}

		fsi.createInstrument(ctx, instruments[i])
	}
}

func (fsi FillingSharesInfo) createInstrument(ctx context.Context, instrumentShare *investapi.Share) {
	instrument := model.NewInstrument(
		instrumentShare.GetFigi(),
		instrumentShare.GetName(),
		instrumentShare.GetSector(),
		instrumentType,
	)

	_, err := fsi.instrumentRepository.Save(ctx, instrument)
	if err != nil {
		fsi.logger.Error(
			log.LogCategoryLogic,
			err.Error(),
			map[string]interface{}{"service": "FillingSharesInfo", "method": "CreateInstrumentsIfNotExist", "action": "save instrument"},
		)
	}

	share := model.NewShare(
		instrumentShare.GetFigi(),
		instrumentShare.GetTicker(),
		instrumentShare.GetClassCode(),
		instrumentShare.GetIsin(),
		instrumentShare.GetLot(),
		instrumentShare.GetCurrency(),
		instrumentShare.GetName(),
		instrumentShare.GetExchange(),
		instrumentShare.GetSector(),
		instrumentShare.GetMinPriceIncrement().GetUnits(),
		instrumentShare.GetMinPriceIncrement().GetNano(),
		instrumentShare.GetApiTradeAvailableFlag(),
	)

	_, err = fsi.shareRepository.Save(ctx, share)
	if err != nil {
		fsi.logger.Error(
			log.LogCategoryLogic,
			err.Error(),
			map[string]interface{}{"service": "FillingSharesInfo", "method": "CreateInstrumentsIfNotExist", "action": "save share"},
		)
	}
}
