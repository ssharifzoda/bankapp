package payment

import (
	"github.com/ssharifzoda/bank/pkg/bank/types"
	"sort"
)

func Max(payments []types.Payment) types.Payment {
	sort.SliceStable(payments, func(i, j int) bool {
		return payments[i].Amount > payments[j].Amount
	})
	return payments[0]
}
