package fillingcurrenciesinfo

import "context"

type FillingCurrenciesInfoInterface interface {
	CheckExistAndLoadInfo(ctx context.Context)
}
