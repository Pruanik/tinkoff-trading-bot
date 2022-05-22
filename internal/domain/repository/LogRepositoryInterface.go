package repository

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type LogRepositoryInterface interface {
	GetLogsDesc(ctx context.Context, lastLogId int, limit int) ([]model.Log, error)
}
