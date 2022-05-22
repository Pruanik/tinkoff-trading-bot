package api

import (
	"net/http"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/gin-gonic/gin"
)

func NewInstrumentApiHandler(
	httpResponseBuilder builder.HttpResponseBuilderInterface,
	instrumentRepository repository.InstrumentRepositoryInterface,
	getInstrumentsBodyBuilder builder.GetInstrumentsBodyBuilderInterface,
) *InstrumentApiHandler {
	return &InstrumentApiHandler{
		httpResponseBuilder:       httpResponseBuilder,
		instrumentRepository:      instrumentRepository,
		getInstrumentsBodyBuilder: getInstrumentsBodyBuilder,
	}
}

type InstrumentApiHandler struct {
	httpResponseBuilder       builder.HttpResponseBuilderInterface
	instrumentRepository      repository.InstrumentRepositoryInterface
	getInstrumentsBodyBuilder builder.GetInstrumentsBodyBuilderInterface
}

func (iah InstrumentApiHandler) Handle(ctx *gin.Context) {
	instruments, err := iah.instrumentRepository.GetInstruments(ctx)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(err.Error()))
	}

	responseBody := iah.getInstrumentsBodyBuilder.CreateBody(instruments)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, iah.httpResponseBuilder.BuildErrorResponse(err.Error()))
	}

	ctx.JSONP(http.StatusOK, iah.httpResponseBuilder.BuildSuccessResponse(responseBody))
}
