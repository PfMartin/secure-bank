package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

const connString = "host=localhost user=root password=secret dbname=secure_bank sslmode=disable"

var testQueries *Queries

func TestMain(m *testing.M) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	defer conn.Close(ctx)


	testQueries = New(conn)

	os.Exit(m.Run())
}