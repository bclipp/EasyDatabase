/*
export DATABASE=testing1234
export USERNAME=testing1234
export PASSWORD=testing1234
export DB_IP_ADDRESS=127.0.0.1
export INT_TEST=true
DATABASE=testing1234;PASSWORD=testing1234;DB_IP_ADDRESS=testing1234;INT_TEST=true;USERNAME=root
*/
package EasyDatabase

import (
	"fmt"
	"os"
	"testing"
)


func TestPostgreSQL_ConnectDisconnect(t *testing.T){
	config := GetVariables()
	var pg = PostgreSQL{
		IPAddress:        config["postgresIP"],
		PostgresPassword: config["postgresPassword"],
		PostgresUser:     config["postgresUser"],
		PostgresDB:       config["postgresDB"],
	}
	err := pg.Connect();if err != nil {
		fmt.Print(err.Error())
	}
	defer pg.Disconnect()
}

func TestPostgreSQL_SendQuery(t *testing.T){
	config := GetVariables()
	var pg = PostgreSQL{
		IPAddress:        config["postgresIP"],
		PostgresPassword: config["postgresPassword"],
		PostgresUser:     config["postgresUser"],
		PostgresDB:       config["postgresDB"],
	}
	err := pg.Connect();if err != nil {
		fmt.Print(err.Error())
	}
	defer pg.Disconnect()
	_,err = pg.SendQuery("SELECT * FROM customers;");if err != nil {
		fmt.Print(err.Error())
	}
}

func TestPostgreSQL_sendQueryReturnData(t *testing.T){
	//config := GetVariables()
	var pg = PostgreSQL{
		IPAddress:        "127.0.0.1",
		PostgresPassword: "testing1234",
		PostgresUser:     "root",
		PostgresDB:       "testing1234",
	}
	err := pg.Connect();if err != nil {
		fmt.Print(err.Error())
	}

	defer pg.Disconnect()
	rows, err := pg.sendQueryReturnData("SELECT latitude,longitude FROM customers;");if err != nil {
		fmt.Print(err.Error())
	}
	defer rows.Close()

	var table []struct {
		Latitude  float64
		Longitude float64
	}

	for rows.Next() {
		var latitude float64
		var longitude float64
		err := rows.Scan(&latitude, &longitude);if err != nil {
			print(err)
		}
		newRow := struct {
			Latitude  float64
			Longitude float64
		}{
			Latitude:  latitude,
			Longitude: longitude,
		}
		table = append(table, newRow)
		fmt.Println(latitude, longitude)
	}
	err = rows.Err();if err != nil {
		print(err)
	}
	err = rows.Close();if err != nil {
		print(err)
	}

}

// get_variables are used to hold environmental variables read by the app
func GetVariables() map[string]string {
	config := make(map[string]string)
	config["postgresDb"] = os.Getenv("DATABASE")
	config["postgresUser"] = os.Getenv("USERNAME")
	config["postgresPassword"] = os.Getenv("PASSWORD")
	config["postgresIP"] = os.Getenv("POSTGRES_IP_ADDRESS")
	config["integrationTest"] = os.Getenv("INT_TEST")

	return config
}
