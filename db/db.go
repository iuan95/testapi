package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "admin"
    password = "admin"
    dbname   = "testapi"
)

type Queries struct {
	DB *pgxpool.Pool
}
var QueriesDb Queries


func DbConnection() error{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	dbpool, err := pgxpool.New(context.Background(), psqlInfo)
	if err != nil {
		return err
	}
	QueriesDb = Queries{
		DB: dbpool,
	}
	defer dbpool.Close()
	return nil
}