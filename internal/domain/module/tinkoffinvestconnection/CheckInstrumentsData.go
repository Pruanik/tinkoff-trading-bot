package tinkoffinvestconnection

import (
	"context"
	"fmt"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
)

func NewCheckInstrumentsData(
	instrumentRepository repository.InstrumentRepositoryInterface,
	instrumentService tinkoffinvest.InstrumentServiceInterface,
) *CheckInstrumentsData {
	return &CheckInstrumentsData{
		instrumentRepository: instrumentRepository,
		instrumentService:    instrumentService,
	}
}

type CheckInstrumentsData struct {
	instrumentRepository repository.InstrumentRepositoryInterface
	instrumentService    tinkoffinvest.InstrumentServiceInterface
}

func (cid CheckInstrumentsData) CheckDataExistAndLoad(ctx context.Context) {
	areSharesExist, err := cid.instrumentRepository.AreInstrumentsExistByType(ctx, "share")
	if err != nil {
		fmt.Println("checkDataExistAndLoad Error: " + err.Error())
	}

	if !areSharesExist {
		fmt.Println("checkDataExistAndLoad is empty")
	} else {
		fmt.Println("checkDataExistAndLoad is full")
	}
}
