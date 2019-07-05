package micro_db

import (
	"github.com/krijnrien/microguild/pkg/messages"
)

type IPriceDatabase interface {
	// ListPrices returns a list of Prices
	ListPrices() ([]*messages.Prices, error)

	//TODO description
	GetDistinctPriceHistoryIds() ([]int, error)

	// AddPrice saves a given Price, ID already assigned.
	AddPrice(b *messages.Prices) (id int64, err error)

	//TODO description
	DropPriceTable() (err error)
}
