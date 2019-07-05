package micro_db

import (
	"github.com/krijnrien/microguild/pkg/messages"
)

type IEventDatabase interface {
	// CreateItem saves a given Item, ID already assigned.
	CreateEvent(b *messages.Event) (err error)
}
