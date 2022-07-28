package main

import (
	"database/sql"
	"log"

	"github.com/WeslleyRibeiro-1999/Crud-Golang/api"
	db "github.com/WeslleyRibeiro-1999/Crud-Golang/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@localhost:5432/go_produtos?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Não foi possivel conectar: ", err)
	}

	store := db.ExecuteNewStore(conn)
	server := api.InstanceServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Api iniciada com erro: ", err)
	}
}
