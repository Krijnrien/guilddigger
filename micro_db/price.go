package micro_db

import (
	"database/sql"
	"fmt"
	"github.com/krijnrien/microguild/pkg/messages"
)

var createPriceTableStatements = []string{
	`CREATE TABLE IF NOT EXISTS events (id string PRIMARY KEY, eventtype string, aggregateid string, aggregatetype string, eventdata string, channel string)`,
}

// itemDatabase persists Item to a MySQL instance.
type priceDatabase struct {
	list                   *sql.Stmt
	distinctPiceHistoryIds *sql.Stmt
	insert                 *sql.Stmt
	dropTable              *sql.Stmt
}

func (db *MySQLConn) PreparePriceStatements() error {
	var err error

	// Prepared statements. The actual SQL queries are in the code near the
	if db.Price.list, err = db.Conn.Prepare(listPriceStatement); err != nil {
		return fmt.Errorf("mysql: prepare list: %v", err)
	}
	if db.Price.distinctPiceHistoryIds, err = db.Conn.Prepare(getDistinctPriceHistoryIdsStatement); err != nil {
		return fmt.Errorf("mysql: prepare get distinct price history ids: %v", err)
	}
	if db.Price.insert, err = db.Conn.Prepare(InsertPriceStatement); err != nil {
		return fmt.Errorf("mysql: prepare insert: %v", err)
	}
	if db.Price.insert, err = db.Conn.Prepare(InsertPriceNowStatement); err != nil {
		return fmt.Errorf("mysql: prepare insert: %v", err)
	}
	if db.MaxAllowedPacket, err = db.Conn.Prepare(getMaxId); err != nil {
		return fmt.Errorf("mysql: get max id: %v", err)
	}
	if db.Price.dropTable, err = db.Conn.Prepare(dropPriceTableStatement); err != nil {
		return fmt.Errorf("mysql: prepare drop: %v", err)
	}

	return nil
}

// Ensure itemDatabase conforms to the IItemDatabase interface.
var _ IPriceDatabase = &priceDatabase{}

const listPriceStatement = `SELECT * FROM price`

// ListItems returns a list of Items, ordered by title.
func (db *priceDatabase) ListPrices() ([]*messages.Prices, error) {
	rows, err := db.list.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Items []*messages.Prices
	for rows.Next() {
		Item, err := scanPrice(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}

		Items = append(Items, Item)
	}

	return Items, nil
}

const getDistinctPriceHistoryIdsStatement = `SELECT DISTINCT itemid FROM price`

// ListItems returns a list of Items, ordered by title.
func (db *priceDatabase) GetDistinctPriceHistoryIds() ([]int, error) {
	rows, err := db.distinctPiceHistoryIds.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Items []int
	for rows.Next() {
		var (
			itemid int
		)
		if err := rows.Scan(&itemid); err != nil {
			return nil, err
		}

		Items = append(Items, itemid)
	}
	return Items, nil
}

const InsertPriceStatement = `INSERT INTO price (itemid, fetched_datetime, buys_quantity, buys_unit_price, sells_quantity, sells_unit_price) VALUES (?, ?, ?, ?, ?, ?)`
const InsertPriceNowStatement = `INSERT IGNORE INTO price (itemid, fetched_datetime, buys_quantity, buys_unit_price, sells_quantity, sells_unit_price) VALUES (?, now(), ?, ?, ?, ?)`

// CreateItem saves a given Item, assigning it a new ID.
func (db *priceDatabase) AddPrice(b *messages.Prices) (id int64, addErr error) {
	r, addErr := execAffectingOneRow(db.insert, b.Id, b.FetchDatetime, b.Buys.Quantity, b.Buys.UnitPrice, b.Sells.Quantity, b.Sells.UnitPrice)
	if addErr != nil {
		return 0, addErr
	}

	lastInsertID, addErr := r.LastInsertId()
	if addErr != nil {
		return 0, fmt.Errorf("mysql: could not get last insert ID: %v", addErr)
	}
	return lastInsertID, nil
}

const getMaxId = `SELECT MAX(itemid) as max_itemid from price;`

// Returns Database max_allowed_packet variable in bytes
func (db *MySQLConn) GetMaxId() (int, error) {
	row := db.MaxAllowedPacket.QueryRow()

	var (
		column string
		value  int
	)

	if err := row.Scan(&column, &value); err != nil {
		return 0, err
	}

	return value, nil
}

const dropPriceTableStatement = `DROP TABLE price`

// CreateItem saves a given Item, assigning it a new ID.
func (db *priceDatabase) DropPriceTable() (err error) {
	_, dropTableErr := execAffectingOneRow(db.dropTable)
	return dropTableErr
}

// scanItem reads a Item from a sql.Row or sql.Rows
func scanPrice(s rowScanner) (*messages.Prices, error) {
	var (
		itemid           int
		fetched_datetime sql.NullString
		buys_quantity    int
		buys_unit_price  int
		sells_quantity   int
		sells_unit_price int
	)
	if err := s.Scan(&itemid, &fetched_datetime, &buys_quantity, &buys_unit_price, &sells_quantity, &sells_unit_price); err != nil {
		return nil, err
	}

	Item := &messages.Prices{
		//Id:             itemid,
		//Whitelisted: whitelisted,
		//Fetch_datetime: fetched_datetime.String,
		//Buys: gw2api.Price{
		//	Quantity:  buys_quantity,
		//	UnitPrice: buys_unit_price,
		//},
		//Sells: gw2api.Price{
		//	Quantity:  sells_quantity,
		//	UnitPrice: sells_unit_price,
		//},
	}
	return Item, nil
}
