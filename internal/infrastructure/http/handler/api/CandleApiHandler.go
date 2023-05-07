package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/gin-gonic/gin"
)

func NewCandleApiHandler(
	httpResponseBuilder builder.HttpResponseBuilderInterface,
	candleRepository repository.CandleRepositoryInterface,
	getCandlesChartBodyBuilder builder.GetCandlesChartBodyBuilderInterface,
) *CandleApiHandler {
	return &CandleApiHandler{
		httpResponseBuilder:        httpResponseBuilder,
		candleRepository:           candleRepository,
		getCandlesChartBodyBuilder: getCandlesChartBodyBuilder,
	}
}

type CandleApiHandler struct {
	httpResponseBuilder        builder.HttpResponseBuilderInterface
	candleRepository           repository.CandleRepositoryInterface
	getCandlesChartBodyBuilder builder.GetCandlesChartBodyBuilderInterface
}

func (cah CandleApiHandler) HandleGetPeriodCandles(ctx *gin.Context) {
	period, figi, err := cah.getPeriodCandleParams(ctx)

	if err != nil {
		ctx.JSONP(http.StatusBadRequest, cah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	fromTime, err := cah.getFromTime(*period)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, cah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	candles, err := cah.candleRepository.GetCandlesByFigiFromTime(ctx, *figi, *fromTime)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, cah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	responseBody := cah.getCandlesChartBodyBuilder.CreateBody(candles)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, cah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	ctx.JSONP(http.StatusOK, cah.httpResponseBuilder.BuildSuccessResponse(responseBody))
}

func (cah CandleApiHandler) getPeriodCandleParams(ctx *gin.Context) (*string, *string, error) {
	queryPeriod, existPeriod := ctx.GetQuery("period")
	figi, existFigi := ctx.GetQuery("figi")

	if !existFigi {
		return nil, nil, errors.New("Figi param does not exist.")
	}

	period := "7d"
	if existPeriod {
		period = queryPeriod
	}

	return &period, &figi, nil
}

func (cah CandleApiHandler) getFromTime(period string) (*time.Time, error) {
	periodArray := strings.Split(period, "")
	if len(periodArray) != 2 {
		return nil, errors.New("Param period is incorrect")
	}

	countDate, err := strconv.Atoi(periodArray[0])
	if err != nil {
		return nil, errors.New("Param period is incorrect")
	}

	var fromTime time.Time
	if periodArray[1] == "d" {
		fromTime = time.Now().AddDate(0, 0, -countDate)
		return &fromTime, nil
	}
	if periodArray[1] == "m" {
		fromTime = time.Now().AddDate(0, -countDate, 0)
		return &fromTime, nil
	}

	return nil, errors.New("Param period is incorrect")
}

func (cah CandleApiHandler) HandleGetLastCandles(ctx *gin.Context) {
	lastId, figi, err := cah.getLastCandleParams(ctx)

	if err != nil {
		ctx.JSONP(http.StatusBadRequest, cah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	candles, err := cah.candleRepository.GetCandlesByFigiFromLastId(ctx, *figi, *lastId)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, cah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	responseBody := cah.getCandlesChartBodyBuilder.CreateBody(candles)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, cah.httpResponseBuilder.BuildErrorResponse(err.Error()))
		return
	}

	ctx.JSONP(http.StatusOK, cah.httpResponseBuilder.BuildSuccessResponse(responseBody))
}

func (cah CandleApiHandler) getLastCandleParams(ctx *gin.Context) (*int, *string, error) {
	queryLastId, existLastId := ctx.GetQuery("lastId")
	figi, existFigi := ctx.GetQuery("figi")

	if !existFigi {
		return nil, nil, errors.New("Figi param does not exist.")
	}

	lastId := 0
	if existLastId {
		lastId, _ = strconv.Atoi(queryLastId)
	}

	return &lastId, &figi, nil
}
