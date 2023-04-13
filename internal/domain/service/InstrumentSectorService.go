package service

import (
	"context"
	"errors"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
)

func NewInstrumentSectorService(instrumentSectorRepository repository.InstrumentSectorRepositoryInterface) InstrumentSectorServiceInterface {
	return &InstrumentSectorService{instrumentSectorRepository: instrumentSectorRepository}
}

type InstrumentSectorService struct {
	instrumentSectorRepository repository.InstrumentSectorRepositoryInterface
}

func (iss *InstrumentSectorService) Create(ctx context.Context, code string) (*model.InstrumentSector, error) {
	sectorName, err := iss.getSectorName(ctx, code)
	if err != nil {
		sectorName = code
	}

	newInstrumentSector := model.NewInstrumentSector(code, sectorName)
	savedInstrumentSector, err := iss.instrumentSectorRepository.Save(ctx, newInstrumentSector)

	return savedInstrumentSector, err
}

func (iss *InstrumentSectorService) CreateIfNotExist(ctx context.Context, code string) (*model.InstrumentSector, error) {
	instrumentSector, err := iss.instrumentSectorRepository.GetInstrumentSectorByCode(ctx, code)

	if instrumentSector == nil {
		instrumentSector, err = iss.Create(ctx, code)
	}

	return instrumentSector, err
}

func (iss *InstrumentSectorService) getSectorName(ctx context.Context, code string) (string, error) {
	codeMap := map[string]string{
		"telecom":         "Telecom",
		"utilities":       "Utilities",
		"green_buildings": "Green Buildings",
		"electrocars":     "Electrocars",
		"industrials":     "Industrials",
		"energy":          "Energy",
		"it":              "IT",
		"financial":       "Financial",
		"health_care":     "Health Care",
		"real_estate":     "Real Estate",
		"other":           "Other",
		"ecomaterials":    "Ecomaterials",
		"currency":        "Ocurrencyne",
		"consumer":        "Consumer",
		"materials":       "Materials",
		"green_energy":    "Green Energy",
	}

	if value, ok := codeMap[code]; ok {
		return value, nil
	} else {
		return "", errors.New("the code doesn't exist in the map")
	}
}
