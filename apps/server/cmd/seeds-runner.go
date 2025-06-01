package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/LaysDragon/blog/apps/server/db/pgrepo/seeds"
	"github.com/LaysDragon/blog/apps/server/internal"
	_ "github.com/lib/pq"
)

// TODO: add config check
type Config struct {
	DBType         string
	DataSourceName string
}

func main() {
	config := internal.LoadConfig()

	db, err := sql.Open(config.DBType, config.DataSourceName)
	if err != nil {
		log.Fatalf("unable to connect to database, %v", err)
	}

	ctx := context.Background()
	seeder := seeds.Seeder{}
	seeder.MinAccountsToSeed = 5
	seeder.MinSitesToSeed = 5
	seeder.MinPostsToSeed = 50

	err = seeder.Run(ctx, db)
	if err != nil {
		panic(err)
	}
}
