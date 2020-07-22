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
	sendQueryReturnData(sqlQuery string, table *[]Row, processRows func(rows *sql.Rows)) ([]struct{}, error)
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
	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		pg.PostgresUser,
		pg.PostgresPassword,
		pg.IPAddress,
		pg.PostgresDB)

	db, err := sql.Open("postgres", connectionString);if err != nil {
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
func (pg *PostgreSQL) sendQueryReturnData(
	sqlQuery string,
	processRows func(rows *sql.Rows) ([]struct{},error))(table []struct{}, err error) {
	rows, err := pg.DB.Query(sqlQuery);if err != nil {
		return nil,err
	}
	defer rows.Close()
	table,err = processRows(rows)
	return table, err
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

/*
// UpdateDbTable is used for taking the the table variable and updating the db
// Params:
//       tableName: the table to query
//return:
//       the error
func (pg *PostgreSQL) UpdateDBTable(table []Row ,tableName string) error {
	for _, row := range table {
		query := UpdateTableQuery(tableName, row)
		result, err := pg.SendQuery(query);if err != nil {
			return err
		}

		count, err := result.RowsAffected();if err != nil {
			return err
		}

		if count != 1 {
			print("Error when updating row, rows effected is not 1.")
		}
	}

	return nil

}
*/