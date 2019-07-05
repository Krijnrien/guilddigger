package persist


type IEventDatabase interface {
	// CreateItem saves a given Item, ID already assigned.
	CreateEvent(b *Event) (err error)
}
