package fillingsharesinfo

import "context"

type FillingSharesInfoInterface interface {
	CheckExistAndLoadInfo(ctx context.Context)
}
