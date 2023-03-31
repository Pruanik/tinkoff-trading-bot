package fillingcurrenciesinfo

import "context"

type FillingCurrenciesInfoInterface interface {
	CreateInstrumentsIfNotExist(ctx context.Context)
}
