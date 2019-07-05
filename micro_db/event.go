package micro_db

import (
	"database/sql"
	"fmt"
	"github.com/krijnrien/microguild/pkg/messages"
)

var createEventTableStatements = []string{`CREATE TABLE IF NOT EXISTS event (id string NOT NULL UNIQUE, eventtype string, aggregateid string, aggregatetype string, eventdata string, channel string, PRIMARY KEY (id))`}

// eventDatabase persists Event to a MySQL instance.
type eventDatabase struct {
	insert *sql.Stmt
}

// Ensure eventDatabase conforms to the IEventDatabase interface.
var _ IEventDatabase = &eventDatabase{}

func (db *MySQLConn) PrepareEventStatements() error {
	var err error

	// Prepared statements. The actual SQL queries are in the code near the
	if db.Event.insert, err = db.Conn.Prepare(InsertEventStatement); err != nil {
		return fmt.Errorf("mysql: prepare insert: %v", err)
	}
	return nil
}

const InsertEventStatement = `INSERT INTO events (id, eventtype, aggregateid, aggregatetype, eventdata, channel) VALUES ($1, $2, $3, $4, $5, $6)`

// CreateEvent saves a given Event, assigning it a new ID.
func (db *eventDatabase) CreateEvent(b *messages.Event) (err error) {
	_, err2 := execAffectingOneRow(db.insert, b.EventId, b.EventType, b.AggregateId, b.AggregateType, b.EventData, b.Channel)
	if err2 != nil {
		return err2
	}
	return nil
}
