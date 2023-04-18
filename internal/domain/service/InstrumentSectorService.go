package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

const defaultSectorCode string = "other"

func NewInstrumentSectorService(
	instrumentSectorRepository repository.InstrumentSectorRepositoryInterface,
	logger log.LoggerInterface,
) InstrumentSectorServiceInterface {
	return &InstrumentSectorService{
		instrumentSectorRepository: instrumentSectorRepository,
		logger:                     logger,
	}
}

type InstrumentSectorService struct {
	instrumentSectorRepository repository.InstrumentSectorRepositoryInterface
	logger                     log.LoggerInterface
}

func (iss *InstrumentSectorService) Create(ctx context.Context, code string) (*model.InstrumentSector, error) {
	sectorName, err := iss.getSectorName(ctx, code)
	if err != nil {
		code = defaultSectorCode
		sectorName, _ = iss.getSectorName(ctx, defaultSectorCode)
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
		message := fmt.Sprintf("the code: `%s` doesn't exist in the codeMap", code)
		iss.logger.Error(
			log.LogCategoryLogic,
			message,
			map[string]interface{}{"service": "InstrumentSectorService", "method": "getSectorName", "action": "convert sector code to name"},
		)

		return "", errors.New(message)
	}
}
