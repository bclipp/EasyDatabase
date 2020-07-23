package EasyDatabase

/*
import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestPostgreSQL_ConnectDisconnect(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	var pg = PostgreSQL{
		IPAddress:        config["postgresIP"],
		PostgresPassword: config["postgresPassword"],
		PostgresUser:     config["postgresUser"],
		PostgresDB:       config["postgresDB"],
	}
	db.
	pg.DB =
	defer pg.Disconnect()
}*/