package service

import (
	"context"

	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

func NewInstrumentService(logger log.LoggerInterface) InstrumentServiceInterface {
	return &InstrumentService{
		logger: logger,
	}
}

type InstrumentService struct {
	logger log.LoggerInterface
}

func (is *InstrumentService) SetInstrumentObservable(ctx context.Context, figi string) error {
	is.logger.Warning(
		log.LogCategoryLogic,
		figi+" set observable",
		map[string]interface{}{"service": "InstrumentService", "method": "SetInstrumentObservable", "action": "set instrument observable"},
	)

	// instrumentSetting := model.NewInstrumentSetting(*figi, *status)
	// _, saveError := iah.instrumentSettingRepository.Update(ctx, instrumentSetting)
	// if saveError != nil {
	// 	ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(saveError.Error()))
	// 	return
	// }
	// iah.fillingHistoricalCandlesData.FillingHistoricalDataByFigi(ctx, *figi)

	return nil
}
