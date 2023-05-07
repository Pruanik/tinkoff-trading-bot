package api

import (
	"net/http"

	"github.com/Pruanik/tinkoff-trading-bot/internal/application/http/api/request"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/fillinghistoricaldata/candles"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/service"
	httpService "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/http/service"
	"github.com/gin-gonic/gin"
)

func NewInstrumentApiHandler(
	httpResponseBuilder builder.HttpResponseBuilderInterface,
	httpService httpService.HttpServiceInterface,
	instrumentRepository repository.InstrumentRepositoryInterface,
	instrumentService service.InstrumentServiceInterface,
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
		httpService:                         httpService,
		instrumentRepository:                instrumentRepository,
		instrumentService:                   instrumentService,
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
	httpService                         httpService.HttpServiceInterface
	instrumentRepository                repository.InstrumentRepositoryInterface
	instrumentService                   service.InstrumentServiceInterface
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
	sectorId := iah.httpService.GetQueryParams(ctx, "sectorId")
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
	sectorId := iah.httpService.GetQueryParams(ctx, "sectorId")
	typeName := iah.httpService.GetQueryParams(ctx, "type")
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

func (iah InstrumentApiHandler) HandleSetInstrumentObservable(ctx *gin.Context) {
	var request request.PostSetInstrumentObservableRequestBody
	ctx.BindJSON(&request)

	if request.Figi == nil {
		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse("Required paramentr 'figi' is absent"))
		return
	}

	error := iah.instrumentService.SetInstrumentObservable(ctx, *request.Figi)

	if error != nil {
		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(error.Error()))
		return
	}

	ctx.JSONP(http.StatusOK, iah.httpResponseBuilder.BuildSuccessResponse(request.Figi))
}

// func (iah InstrumentApiHandler) HandleGetCollectingInstruments(ctx *gin.Context) {
// 	instrumentsSettings, err := iah.instrumentSettingRepository.GetInstrumentsSettingsWhereIsCollectingTrue(ctx)
// 	if err != nil {
// 		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(err.Error()))
// 		return
// 	}

// 	responseBody := iah.getCollectingInstrumentsBodyBuilder.CreateBody(instrumentsSettings)
// 	if err != nil {
// 		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(err.Error()))
// 		return
// 	}

// 	ctx.JSONP(http.StatusOK, iah.httpResponseBuilder.BuildSuccessResponse(responseBody))
// }
