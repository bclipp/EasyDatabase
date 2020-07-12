package easydb

import (
	"fmt"
	"testing"
)

func TestPG(t *testing.T) {
	//CheckIntegrationTest(t)
	t.Run("loadTable", func(t *testing.T) {
		config := GetVariables()
		var pg = PostgreSQL{
			IPAddress:        config["postgresIP"],
			PostgresPassword: config["postgresPassword"],
			PostgresUser:     config["postgresUser"],
			PostgresDB:       config["postgresDB"],
		}

		err := pg.connect();if err != nil {
			fmt.Print(err.Error())
		}
		defer pg.close()
		var table []Row
		fmt.Println("sending query")
		table, err = pg.returnTable("customers", -1)
		if err != nil {fmt.Print(err.Error())}

		if len(table) < 1 {
			t.Errorf("Error, reading customers table and no data was found.")
		}
	})
	
	t.Run("UpdateTable", func(t *testing.T) {
		config := GetVariables()
		var pg = PostgreSQL{
			IPAddress:        config["IpAddress"],
			PostgresPassword: config["postgresPassword"],
			PostgresUser:     config["postgresUser"],
			PostgresDB:       config["postgresDB"],
		}
		err := pg.connect()
		if err != nil {fmt.Print(err.Error())}
		defer pg.close()
	})
}
