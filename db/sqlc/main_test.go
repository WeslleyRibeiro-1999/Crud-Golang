package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/go_produtos?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Não foi possivel conectar: ", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
