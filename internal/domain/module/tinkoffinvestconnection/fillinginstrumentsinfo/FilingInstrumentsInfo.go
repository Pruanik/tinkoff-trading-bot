package fillinginstrumentsinfo

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/fillinginstrumentsinfo/fillingcurrenciesinfo"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/fillinginstrumentsinfo/fillingsharesinfo"
)

func NewFillingInstrumentsInfo(
	fillingSharesInfo fillingsharesinfo.FillingSharesInfoInterface,
	fillingCurrenciesInfo fillingcurrenciesinfo.FillingCurrenciesInfoInterface,
) FillingInstrumentsInfoInterface {
	return &FillingInstrumentsInfo{
		fillingSharesInfo:     fillingSharesInfo,
		fillingCurrenciesInfo: fillingCurrenciesInfo,
	}
}

type FillingInstrumentsInfo struct {
	fillingSharesInfo     fillingsharesinfo.FillingSharesInfoInterface
	fillingCurrenciesInfo fillingcurrenciesinfo.FillingCurrenciesInfoInterface
}

func (fii FillingInstrumentsInfo) LoadInfo(ctx context.Context) {
	fii.fillingSharesInfo.CheckExistAndLoadInfo(ctx)
	fii.fillingCurrenciesInfo.CheckExistAndLoadInfo(ctx)
}
