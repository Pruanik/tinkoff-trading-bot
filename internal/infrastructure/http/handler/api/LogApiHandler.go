package api

import (
	"net/http"
	"strconv"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/gin-gonic/gin"
)

func NewLogApiHandler(
	httpResponseBuilder builder.HttpResponseBuilderInterface,
	logRepository repository.LogRepositoryInterface,
	getLogsBodyBuilder builder.GetLogsBodyBuilderInterface,
) *LogApiHandler {
	return &LogApiHandler{
		httpResponseBuilder: httpResponseBuilder,
		logRepository:       logRepository,
		getLogsBodyBuilder:  getLogsBodyBuilder,
	}
}

type LogApiHandler struct {
	httpResponseBuilder builder.HttpResponseBuilderInterface
	logRepository       repository.LogRepositoryInterface
	getLogsBodyBuilder  builder.GetLogsBodyBuilderInterface
}

func (lah LogApiHandler) Handle(ctx *gin.Context) {
	lastId, limit := lah.getLogsParams(ctx)

	logs, err := lah.logRepository.GetLogsDesc(ctx, lastId, limit)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, lah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	responseBody := lah.getLogsBodyBuilder.CreateBody(logs)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, lah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	ctx.JSONP(http.StatusOK, lah.httpResponseBuilder.BuildSuccessResponse(responseBody))
}

func (lah LogApiHandler) getLogsParams(ctx *gin.Context) (int, int) {
	queryLastId, existLastId := ctx.GetQuery("lastId")
	queryLimit, existLimit := ctx.GetQuery("limit")

	lastId := 0
	if existLastId {
		lastId, _ = strconv.Atoi(queryLastId)
	}

	limit := 100
	if existLimit {
		limit, _ = strconv.Atoi(queryLimit)
	}

	return lastId, limit
}
