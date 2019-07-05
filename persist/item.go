package persist

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

var createItemTableStatements = []string{
	`CREATE TABLE IF NOT EXISTS item (item_id int4 PRIMARY KEY, item_name string, icon string, description string, item_type string, rarity string, item_level int4, vendorvalue int4)`,
}

// itemDatabase persists Item to a MySQL instance.
type itemDatabase struct {
	list       *sql.Stmt
	listIds    *sql.Stmt
	bulkInsert *sql.DB
	insert     *sql.Stmt
	get        *sql.Stmt
	update     *sql.Stmt
	delete     *sql.Stmt
}

// Ensure itemDatabase conforms to the IItemDatabase interface.
var _ IItemDatabase = &itemDatabase{}

func (db *MySQLConn) PrepareItemStatements() error {
	var err error

	// Prepared statements. The actual SQL queries are in the code near the
	if db.Item.list, err = db.Conn.Prepare(listItemsStatement); err != nil {
		return fmt.Errorf("mysql: prepare list: %v", err)
	}
	if db.Item.listIds, err = db.Conn.Prepare(listItemIdsStatement); err != nil {
		return fmt.Errorf("mysql: prepare list Ids: %v", err)
	}
	if db.Item.get, err = db.Conn.Prepare(getItemStatement); err != nil {
		return fmt.Errorf("mysql: prepare get: %v", err)
	}
	if db.Item.insert, err = db.Conn.Prepare(InsertItemStatement); err != nil {
		return fmt.Errorf("mysql: prepare insert: %v", err)
	}
	if db.Item.update, err = db.Conn.Prepare(updateItemStatement); err != nil {
		return fmt.Errorf("mysql: prepare update: %v", err)
	}
	if db.Item.delete, err = db.Conn.Prepare(deleteItemStatement); err != nil {
		return fmt.Errorf("mysql: prepare delete: %v", err)
	}
	return nil
}

const listItemsStatement = `SELECT * FROM item`

// ListItems returns a list of Items, ordered by title.
func (db *itemDatabase) ListItems() ([]*Item, error) {
	rows, err := db.list.Query(listItemsStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Items []*Item
	for rows.Next() {
		Item, err := scanItem(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}

		Items = append(Items, Item)
	}

	return Items, nil
}

const listItemIdsStatement = `SELECT item_id FROM item`

//read item ids
func (db *itemDatabase) ListItemIds() ([]*ItemResponse, error) {
	rows, err := db.listIds.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse all rows into an array of ItemResponse
	var ItemResponses []*ItemResponse
	for rows.Next() {
		itemResponse := ItemResponse{}
		if err = rows.Scan(&itemResponse.ItemId); err == nil {
			ItemResponses = append(ItemResponses, &itemResponse)
		}
	}
	if err != nil {
		return nil, errors.Wrap(err, "error on reading all items")
	}
	return ItemResponses, nil
}

const getItemStatement = "SELECT * FROM item WHERE id = ?"

// GetItem retrieves a Item by its ID.
func (db *itemDatabase) GetItem(id int) (*Item, error) {
	Item, err := scanItem(db.get.QueryRow(id))
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("mysql: could not find Item with id %d", id)
	}
	if err != nil {
		return nil, fmt.Errorf("mysql: could not get Item: %v", err)
	}
	return Item, nil
}

//TODO Ignoring duplication error for now, check before inserting if exists or values changes?
const InsertItemStatement = `INSERT IGNORE INTO item (id, item_name, icon, description, item_type, rarity, item_level, vendorValue) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

// CreateItem saves a given Item, assigning it a new ID.
func (db *itemDatabase) CreateItem(b *Item) (id int64, err error) {
	r, err := execAffectingOneRow(db.insert, b.ID, b.Name, b.Description, b.Type, b.Level, b.Rarity, b.VendorValue, b.Icon)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := r.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("mysql: could not get last insert ID: %v", err)
	}
	return lastInsertID, nil
}

func (db *MySQLConn) BulkCreateItem(unsavedRows []Item) error {
	valueStrings := make([]string, 0, len(unsavedRows))
	valueArgs := make([]interface{}, 0, len(unsavedRows)*8)
	i := 0
	for _, post := range unsavedRows {
		i++
		var placeholderString []int
		baseI := i * 8
		placeholderString = append(placeholderString, baseI-7, baseI-6, baseI-5, baseI-4, baseI-3, baseI-2, baseI-1, baseI)
		valueStrings = append(valueStrings, fmt.Sprintf("($%s)", strings.Trim(strings.Replace(fmt.Sprint(placeholderString), " ", ", $", -1), "[]")))
		valueArgs = append(valueArgs, post.ID)
		valueArgs = append(valueArgs, post.Name)
		valueArgs = append(valueArgs, post.Icon)
		valueArgs = append(valueArgs, post.Description)
		valueArgs = append(valueArgs, post.Type)
		valueArgs = append(valueArgs, post.Rarity)
		valueArgs = append(valueArgs, post.Level)
		valueArgs = append(valueArgs, post.VendorValue)
	}
	stmt := fmt.Sprintf("INSERT INTO item (item_id, item_name, icon, description, item_type, rarity, item_level, vendorvalue) VALUES %s", strings.Join(valueStrings, ","))
	_, err := db.Conn.Exec(stmt, valueArgs...)
	if err != nil {
		return err
	}
	return nil
}

const deleteItemStatement = `DELETE FROM item WHERE id = ?`

// DeleteItem removes a given Item by its ID.
func (db *itemDatabase) DeleteItem(id int64) error {
	if id == 0 {
		return errors.New("mysql: Item with unassigned ID passed into deleteItem")
	}
	_, err := execAffectingOneRow(db.delete, id)
	return err
}

const updateItemStatement = `UPDATE item SET item_id=?, item_name=?, icon=?, description=?, item_type=?, rarity=?, item_level=?, vendorValue=?, WHERE id = ?`

// UpdateItem updates the entry for a given Item.
func (db *itemDatabase) UpdateItem(b *Item) error {
	if b.ID == 0 {
		return errors.New("mysql: Item with unassigned ID passed into updateItem")
	}

	_, err := execAffectingOneRow(db.update, b.ID, b.Name, b.Description,
		b.Type, b.Rarity, b.VendorValue, b.Icon)
	return err
}

// scanItem reads a Item from a sql.Row or sql.Rows
func scanItem(s rowScanner) (*Item, error) {
	var (
		id          uint32
		name        sql.NullString
		description sql.NullString
		itemType    sql.NullString
		level       uint32
		rarity      sql.NullString
		vendorValue uint32
		icon        sql.NullString
	)
	if err := s.Scan(&id, &name, &description, &itemType, &level, &rarity, &vendorValue, &icon); err != nil {
		return nil, err
	}

	Item := &Item{
		ID:          id,
		Name:        name.String,
		Description: description.String,
		Type:        itemType.String,
		Level:       level,
		Rarity:      description.String,
		VendorValue: vendorValue,
		Icon:        icon.String,
	}
	return Item, nil
}
