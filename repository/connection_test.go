package repository

import (
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/go-hclog"
)

func TestConnection(t *testing.T) {
	logger := hclog.New(
		&hclog.LoggerOptions{
			Name:  "testing-database",
			Level: hclog.LevelFromString("DEBUG"),
		},
	)
	connection := "root:@tcp(localhost:3306)/golang-db"
	conn := NewConnection(logger, connection, "mysql", 20, 5, time.Minute, time.Minute*30)
	db, err := conn.CreateConnection()
	defer db.Close()
	if err != nil {
		t.Fatal(err.Error())
	}
	ping := db.Ping()
	if ping != nil {
		t.Fatal(ping.Error())
	}
	fmt.Println("Success without any error")
}
