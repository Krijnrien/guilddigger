package micro_db

import (
	"github.com/krijnrien/microguild/pkg/messages"
)

type IItemDatabase interface {
	// ListItems returns a list of items
	ListItems() ([]*messages.Item, error)

	// GetItem retrieves a book by its ID.
	GetItem(id int) (*messages.Item, error)

	// CreateItem saves a given Item, ID already assigned.
	CreateItem(b *messages.Item) (id int64, err error)

	// DeleteItem removes a given Item by its ID.
	DeleteItem(id int64) error

	// UpdateItem updates the entry for a given book.
	UpdateItem(b *messages.Item) error
}
