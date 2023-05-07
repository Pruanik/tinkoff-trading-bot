package service

import (
	"context"
)

type InstrumentServiceInterface interface {
	SetInstrumentObservable(ctx context.Context, figi string) error
}
