//This module is used for interacting with the postgresql database

package EasyDatabase

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Row struct {
	BlockID   int
	StateCode string
	StateFips int
	BlockPop  int
	ID        int
	Latitude  float64
	Longitude float64
}

type Database interface {
	connect()  error
	Disconnect()
	sendQueryReturnData(sqlQuery string)(*sql.Rows, error)
	sendQuery(query string)  (sql.Result, error)
}

// Database is used to hold the connection related variables
type PostgreSQL struct {
	DB               *sql.DB
	IPAddress        string
	PostgresPassword string
	PostgresUser     string
	PostgresDB       string
}

// Connect is used to handle connecting to the database
// Params:
// return:
//       error from the connection setup
func (pg *PostgreSQL) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pg.IPAddress, pg.PostgresUser, pg.PostgresPassword, pg.PostgresDB)

	db, err := sql.Open("postgres", psqlInfo);if err != nil {
		return err
	}
	pg.DB = db

	err = pg.DB.Ping();if err != nil {
		return err
	}

	return nil
}

// Close is used to handle closing the connection to the database
// Params:
// return:
//       error from the connection setup
func (pg *PostgreSQL) Disconnect() {
	_ = pg.DB.Close()
}

// ReadTable is used for reading data from the database and storing it in the
// table field
// Params:
//       tableName: the table to query
//return:
//       Jason return document
//       rest http response code
//       the error
func (pg *PostgreSQL) sendQueryReturnData(sqlQuery string)(*sql.Rows, error) {
	rows, err := pg.DB.Query(sqlQuery)
	return rows, err
}

// SendQuery is used for sending query to a database
// Params:
//       SendQuery: SQL to send
//return:
//		 result variable , see result interface doc in sql
//       the error
func (pg *PostgreSQL) SendQuery(query string) (sql.Result, error) {
	result, err := pg.DB.Exec(query)
	if err != nil { return result, err}

	return result, nil
}
