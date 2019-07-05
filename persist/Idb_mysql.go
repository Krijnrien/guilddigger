package persist

type IdbMysql interface {
	// Close closes the database, freeing up any available resources.
	Close() error
}
