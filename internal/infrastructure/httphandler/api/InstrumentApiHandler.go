package api

import (
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
	getInstrumentTypesBodyBuilder builder.GetInstrumentTypesBodyBuilderInterface,
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
		getInstrumentTypesBodyBuilder:       getInstrumentTypesBodyBuilder,
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
	getInstrumentTypesBodyBuilder       builder.GetInstrumentTypesBodyBuilderInterface
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

func (iah InstrumentApiHandler) HandleGetTypes(ctx *gin.Context) {
	sectorId := iah.getQueryParams(ctx, "sectorId")
	var instrumentTypes []string
	var err error

	if sectorId != nil {
		instrumentTypes, err = iah.instrumentRepository.GetInstrumentTypesBySectorId(ctx, *sectorId)
	} else {
		instrumentTypes, err = iah.instrumentRepository.GetInstrumentTypes(ctx)
	}

	if err != nil {
		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	responseBody := iah.getInstrumentTypesBodyBuilder.CreateBody(instrumentTypes)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	ctx.JSONP(http.StatusOK, iah.httpResponseBuilder.BuildSuccessResponse(responseBody))
}

func (iah InstrumentApiHandler) HandleGetInstruments(ctx *gin.Context) {
	sectorId := iah.getQueryParams(ctx, "sectorId")
	typeName := iah.getQueryParams(ctx, "type")
	var instruments []model.Instrument
	var err error

	if sectorId != nil && typeName != nil {
		instruments, err = iah.instrumentRepository.GetInstrumentsBySectorIdAndType(ctx, *sectorId, *typeName)
	} else if sectorId != nil {
		instruments, err = iah.instrumentRepository.GetInstrumentsBySectorId(ctx, *sectorId)
	} else if typeName != nil {
		instruments, err = iah.instrumentRepository.GetInstrumentsByType(ctx, *typeName)
	} else {
		instruments, err = iah.instrumentRepository.GetInstruments(ctx)
	}

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

// func (iah InstrumentApiHandler) HandleSetCollectingInstrument(ctx *gin.Context) {
// 	figi, status, err := iah.getSetCollectingInstrumentParams(ctx)

// 	if err != nil {
// 		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(err.Error()))
// 		return
// 	}

// 	instrumentSetting := model.NewInstrumentSetting(*figi, *status)
// 	_, saveError := iah.instrumentSettingRepository.Update(ctx, instrumentSetting)
// 	if saveError != nil {
// 		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(saveError.Error()))
// 		return
// 	}
// 	iah.fillingHistoricalCandlesData.FillingHistoricalDataByFigi(ctx, *figi)

// 	ctx.JSONP(http.StatusOK, iah.httpResponseBuilder.BuildSuccessResponse(nil))
// }

func (iah InstrumentApiHandler) getQueryParams(ctx *gin.Context, paramKey string) *string {
	paramValue, isParamExist := ctx.GetQuery(paramKey)

	var result *string
	if isParamExist {
		result = &paramValue
	}

	return result
}
