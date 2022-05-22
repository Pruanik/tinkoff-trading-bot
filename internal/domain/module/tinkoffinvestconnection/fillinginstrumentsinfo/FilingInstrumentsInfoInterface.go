package fillinginstrumentsinfo

import "context"

type FillingInstrumentsInfoInterface interface {
	LoadInfo(ctx context.Context)
}
