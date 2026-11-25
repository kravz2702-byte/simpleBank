package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbSource = "postgresql://postgres:qwertyvbnm@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testConn *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	ctx := context.Background()
	testConn, err = pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testConn)

	os.Exit(m.Run())
}
