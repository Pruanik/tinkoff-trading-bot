package fillingsharesinfo

import "context"

type FillingSharesInfoInterface interface {
	CreateInstrumentsIfNotExist(ctx context.Context)
}
