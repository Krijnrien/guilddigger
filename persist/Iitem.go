package persist

type IItemDatabase interface {
	// ListItems returns a list of items
	ListItems() ([]*Item, error)

	// GetItem retrieves a book by its ID.
	GetItem(id int) (*Item, error)

	// CreateItem saves a given Item, ID already assigned.
	CreateItem(b *Item) (id int64, err error)

	// DeleteItem removes a given Item by its ID.
	DeleteItem(id int64) error

	// UpdateItem updates the entry for a given book.
	UpdateItem(b *Item) error
}
