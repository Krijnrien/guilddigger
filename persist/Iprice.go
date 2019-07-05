package persist

type IPriceDatabase interface {
	// ListPrices returns a list of Prices
	ListPrices() ([]*Prices, error)

	//TODO description
	GetDistinctPriceHistoryIds() ([]int, error)

	// AddPrice saves a given Price, ID already assigned.
	AddPrice(b *Prices) (id int64, err error)

	//TODO description
	DropPriceTable() (err error)
}
