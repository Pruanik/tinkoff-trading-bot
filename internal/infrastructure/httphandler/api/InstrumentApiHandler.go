package api

import (
	"errors"
	"net/http"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/fillinghistoricaldata/candles"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/gin-gonic/gin"
)

func NewInstrumentApiHandler(
	httpResponseBuilder builder.HttpResponseBuilderInterface,
	instrumentRepository repository.InstrumentRepositoryInterface,
	instrumentSectorRepository repository.InstrumentSectorRepositoryInterface,
	instrumentSettingRepository repository.InstrumentSettingRepositoryInterface,
	getInstrumentsBodyBuilder builder.GetInstrumentsBodyBuilderInterface,
	getInstrumentSectorsBodyBuilder builder.GetInstrumentSectorsBodyBuilderInterface,
	getCollectingInstrumentsBodyBuilder builder.GetCollectingInstrumentsBodyBuilderInterface,
	fillingHistoricalCandlesData candles.FillingHistoricalCandlesDataInterface,
) *InstrumentApiHandler {
	return &InstrumentApiHandler{
		httpResponseBuilder:                 httpResponseBuilder,
		instrumentRepository:                instrumentRepository,
		instrumentSectorRepository:          instrumentSectorRepository,
		instrumentSettingRepository:         instrumentSettingRepository,
		getInstrumentsBodyBuilder:           getInstrumentsBodyBuilder,
		getInstrumentSectorsBodyBuilder:     getInstrumentSectorsBodyBuilder,
		getCollectingInstrumentsBodyBuilder: getCollectingInstrumentsBodyBuilder,
		fillingHistoricalCandlesData:        fillingHistoricalCandlesData,
	}
}

type InstrumentApiHandler struct {
	httpResponseBuilder                 builder.HttpResponseBuilderInterface
	instrumentRepository                repository.InstrumentRepositoryInterface
	instrumentSectorRepository          repository.InstrumentSectorRepositoryInterface
	instrumentSettingRepository         repository.InstrumentSettingRepositoryInterface
	getInstrumentsBodyBuilder           builder.GetInstrumentsBodyBuilderInterface
	getInstrumentSectorsBodyBuilder     builder.GetInstrumentSectorsBodyBuilderInterface
	getCollectingInstrumentsBodyBuilder builder.GetCollectingInstrumentsBodyBuilderInterface
	fillingHistoricalCandlesData        candles.FillingHistoricalCandlesDataInterface
}

func (iah InstrumentApiHandler) HandleGetSectors(ctx *gin.Context) {
	instrumentSectors, err := iah.instrumentSectorRepository.GetSectors(ctx)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	responseBody := iah.getInstrumentSectorsBodyBuilder.CreateBody(instrumentSectors)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	ctx.JSONP(http.StatusOK, iah.httpResponseBuilder.BuildSuccessResponse(responseBody))
}

func (iah InstrumentApiHandler) HandleGetInstruments(ctx *gin.Context) {
	instruments, err := iah.instrumentRepository.GetInstruments(ctx)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	responseBody := iah.getInstrumentsBodyBuilder.CreateBody(instruments)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	ctx.JSONP(http.StatusOK, iah.httpResponseBuilder.BuildSuccessResponse(responseBody))
}

func (iah InstrumentApiHandler) HandleGetCollectingInstruments(ctx *gin.Context) {
	instrumentsSettings, err := iah.instrumentSettingRepository.GetInstrumentsSettingsWhereIsCollectingTrue(ctx)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	responseBody := iah.getCollectingInstrumentsBodyBuilder.CreateBody(instrumentsSettings)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	ctx.JSONP(http.StatusOK, iah.httpResponseBuilder.BuildSuccessResponse(responseBody))
}

func (iah InstrumentApiHandler) HandleSetCollectingInstrument(ctx *gin.Context) {
	figi, status, err := iah.getSetCollectingInstrumentParams(ctx)

	if err != nil {
		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	instrumentSetting := model.NewInstrumentSetting(*figi, *status)
	_, saveError := iah.instrumentSettingRepository.Update(ctx, instrumentSetting)
	if saveError != nil {
		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(saveError.Error()))
		return
	}
	iah.fillingHistoricalCandlesData.FillingHistoricalDataByFigi(ctx, *figi)

	ctx.JSONP(http.StatusOK, iah.httpResponseBuilder.BuildSuccessResponse(nil))
}

func (iah InstrumentApiHandler) getSetCollectingInstrumentParams(ctx *gin.Context) (*string, *bool, error) {
	figi, existFigi := ctx.GetQuery("figi")
	statusQuery, existStatus := ctx.GetQuery("status")

	if !existFigi {
		return nil, nil, errors.New("Figi param does not exist.")
	}

	status := false
	if existStatus && statusQuery == "true" {
		status = true
	}

	return &figi, &status, nil
}
