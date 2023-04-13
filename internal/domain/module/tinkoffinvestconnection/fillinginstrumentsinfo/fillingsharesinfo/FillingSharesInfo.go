package fillingsharesinfo

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/service"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc/investapi"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

const instrumentType string = "share"

func NewFillingSharesInfo(
	instrumentSectorService service.InstrumentSectorServiceInterface,
	instrumentRepository repository.InstrumentRepositoryInterface,
	shareRepository repository.ShareRepositoryInterface,
	instrumentService tinkoffinvest.InstrumentServiceInterface,
	logger log.LoggerInterface,
) FillingSharesInfoInterface {
	return &FillingSharesInfo{
		instrumentSectorService: instrumentSectorService,
		instrumentRepository:    instrumentRepository,
		shareRepository:         shareRepository,
		instrumentService:       instrumentService,
		logger:                  logger,
	}
}

type FillingSharesInfo struct {
	instrumentSectorService service.InstrumentSectorServiceInterface
	instrumentRepository    repository.InstrumentRepositoryInterface
	shareRepository         repository.ShareRepositoryInterface
	instrumentService       tinkoffinvest.InstrumentServiceInterface
	logger                  log.LoggerInterface
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

		fsi.createInstrumentShare(ctx, instruments[i])
	}
}

func (fsi FillingSharesInfo) createInstrumentShare(ctx context.Context, instrumentShare *investapi.Share) {
	sector, sectorErr := fsi.createSector(ctx, instrumentShare.GetSector())
	if sectorErr != nil {
		fsi.logger.Error(
			log.LogCategoryLogic,
			sectorErr.Error(),
			map[string]interface{}{"service": "FillingSharesInfo", "method": "createSector", "action": "save sector"},
		)
		return
	}

	_, instrumentErr := fsi.createInstrument(ctx, instrumentShare, sector)
	if instrumentErr != nil {
		fsi.logger.Error(
			log.LogCategoryLogic,
			instrumentErr.Error(),
			map[string]interface{}{"service": "FillingSharesInfo", "method": "createInstrument", "action": "save instrument"},
		)
	}

	_, shareErr := fsi.createShare(ctx, instrumentShare)
	if shareErr != nil {
		fsi.logger.Error(
			log.LogCategoryLogic,
			shareErr.Error(),
			map[string]interface{}{"service": "FillingSharesInfo", "method": "createShare", "action": "save share"},
		)
	}
}

func (fsi FillingSharesInfo) createSector(ctx context.Context, code string) (*model.InstrumentSector, error) {
	return fsi.instrumentSectorService.CreateIfNotExist(ctx, code)
}

func (fsi FillingSharesInfo) createInstrument(
	ctx context.Context,
	instrumentShare *investapi.Share,
	instrumentSector *model.InstrumentSector,
) (*model.Instrument, error) {
	newInstrument := model.NewInstrument(
		instrumentShare.GetFigi(),
		instrumentShare.GetName(),
		instrumentSector.Id,
		instrumentType,
	)
	return fsi.instrumentRepository.Save(ctx, newInstrument)
}

func (fsi FillingSharesInfo) createShare(ctx context.Context, instrumentShare *investapi.Share) (*model.Share, error) {
	newShare := model.NewShare(
		instrumentShare.GetFigi(),
		instrumentShare.GetTicker(),
		instrumentShare.GetClassCode(),
		instrumentShare.GetIsin(),
		instrumentShare.GetLot(),
		instrumentShare.GetCurrency(),
		instrumentShare.GetName(),
		instrumentShare.GetExchange(),
		instrumentShare.GetMinPriceIncrement().GetUnits(),
		instrumentShare.GetMinPriceIncrement().GetNano(),
		instrumentShare.GetApiTradeAvailableFlag(),
	)

	return fsi.shareRepository.Save(ctx, newShare)
}
