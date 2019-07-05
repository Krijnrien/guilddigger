package persist

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	_ "github.com/lib/pq"
)

var createAllTablesStatements = [][]string{
	createItemTableStatements,
	createPriceTableStatements,
}

type MySQLConfig struct {
	// Optional.
	Username, Password string

	// Host of the MySQL instance.
	//
	// If set, UnixSocket should be unset.
	Host string

	// Port of the MySQL instance.
	//
	// If set, UnixSocket should be unset.
	Port int

	// UnixSocket is the file path to a unix socket.
	//
	// If set, Host and Port should be unset.
	UnixSocket string
}

type MySQLConn struct {
	Conn             *sql.DB
	MaxAllowedPacket *sql.Stmt
	Item             *itemDatabase
	Price            *priceDatabase
	Event            *eventDatabase
}

// Enforce interface
var _ IdbMysql = &MySQLConn{}

// dataStoreName returns a connection string suitable for sql.Open.
func (mysqlConfig MySQLConfig) dataStoreName(databaseName string) string {
	var cred string
	// [username[:password]@]
	if mysqlConfig.Username != "" {
		cred = mysqlConfig.Username
		if mysqlConfig.Password != "" {
			cred = cred + ":" + mysqlConfig.Password
		}
		cred = cred + "@"
	}

	if mysqlConfig.UnixSocket != "" {
		return fmt.Sprintf("%sunix(%s)/%s", cred, mysqlConfig.UnixSocket, databaseName)
	}
	return fmt.Sprintf("%stcp([%s]:%d)/%s", cred, mysqlConfig.Host, mysqlConfig.Port, databaseName)
}

// newMySQLDB creates a new MySQLConn backed by a given MySQL server.
func (mysqlConfig MySQLConfig) NewMySQLDB() (*MySQLConn, error) {
	// Check database and table exists. If not, create it.
	if err := mysqlConfig.ensureTableExists(); err != nil {
		return nil, err
	}

	//conn, err := sql.Open("mysql", mysqlConfig.dataStoreName("gowars"))
	conn, err := sql.Open("postgres", "postgresql://root@192.168.99.100:26257/ordersdb?sslmode=disable")
	//conn.SetMaxIdleConns(10)
	//conn.SetMaxOpenConns(50)
	if err != nil {
			return nil, fmt.Errorf("mysql: could not get a connection: %v", err)
	}
	if err := conn.Ping(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("mysql: could not establish a good connection: %v", err)
	}

	db := &MySQLConn{
		Conn:  conn,
		Item:  &itemDatabase{},
		Price: &priceDatabase{},
		Event: &eventDatabase{},
	}

	db.PrepareItemStatements()
	db.PreparePriceStatements()
	db.PrepareEventStatements()

	return db, nil
}

// rowScanner is implemented by sql.Row and sql.Rows
type rowScanner interface {
	Scan(dest ...interface{}) error
}

// ensureTableExists checks the table exists. If not, it creates it.
func (mysqlConfig MySQLConfig) ensureTableExists() error {
	conn, err := sql.Open("postgres", "postgresql://root@192.168.99.100:26257/ordersdb?sslmode=disable")
	//conn, err := sql.Open("mysql", mysqlConfig.dataStoreName("gowars"))
	if err != nil {
		return fmt.Errorf("mysql: could not get a connection: %v", err)
	}
	defer conn.Close()

	// Check the connection.
	if conn.Ping() == driver.ErrBadConn {
		return fmt.Errorf("mysql: could not connect to the database. " +
			"could be bad address, or this address is not whitelisted for access")
	}

	return createTable(conn)
}

// createTable creates the table, and if necessary, the database.
func createTable(conn *sql.DB) error {

	if _, err := conn.Exec(
		"CREATE DATABASE IF NOT EXISTS ordersdb; use ordersdb;"); err != nil {
		return err
	}

	for _, createStatements := range createAllTablesStatements {
		for _, createStmt := range createStatements {
			_, err := conn.Exec(createStmt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// execAffectingOneRow executes a given statement, expecting one row to be affected.
func execAffectingOneRow(stmt *sql.Stmt, args ...interface{}) (sql.Result, error) {
	r, err := stmt.Exec(args...)
	if err != nil {
		return r, fmt.Errorf("mysql: could not execute statement: %v", err)
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return r, fmt.Errorf("mysql: could not get rows affected: %v", err)
	} else if rowsAffected != 1 {
		return r, fmt.Errorf("mysql: expected 1 row affected, got %d", rowsAffected)
	}
	return r, nil
}

// Close closes the database, freeing up any resources.
// Before closing it finishing flushing any existing batch inserts.
func (db *MySQLConn) Close() error {
	return db.Conn.Close()
}
