package db

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

	"github.com/Thiht/transactor"
	stdlibTransactor "github.com/Thiht/transactor/stdlib"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("db",
	fx.Provide(
		InitDb,
		func(db *sql.DB) boil.ContextExecutor {
			return db
		},
		func(db *sql.DB) (transactor.Transactor, stdlibTransactor.DBGetter) {
			return stdlibTransactor.NewTransactor(
				db,
				stdlibTransactor.NestedTransactionsSavepoints,
			)
		},
	),
	fx.Invoke(
		InitDbSchema,
	),
)

type DbConfig interface {
	GetDBType() string
	GetDataSourceName() string
}

func InitDb(config DbConfig) (*sql.DB, error) {
	db, err := sql.Open(config.GetDBType(), config.GetDataSourceName())
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database, %w", err)
	}
	err = db.PingContext(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database, %w", err)
	}
	return db, nil
}

//go:embed init.sql
var initSql string

//go:embed model.sql
var modelSql string

func InitDbSchema(db *sql.DB, log *zap.Logger) error {
	log = log.Named("db")
	var count int
	err := db.QueryRow("SELECT count(*) FROM information_schema.tables where table_schema = 'public'").Scan(&count)
	if err != nil {
		log.Fatal("failed to check rule table", zap.Error(err))
	}

	if count == 0 {
		_, err := db.Exec(modelSql)
		if err != nil {
			return err
		}
		log.Info("setup the needed db table completed")
	} else {
		log.Info("no need to init db table")
	}
	return nil

}

func InitDbData(db *sql.DB, log *zap.Logger) error {
	log = log.Named("db")
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM account").Scan(&count)
	if err != nil {
		log.Fatal("failed to check rule table", zap.Error(err))
	}

	if count == 0 {
		_, err = db.Exec(initSql)
		if err != nil {
			return err
		}
		log.Info("setup the needed init db data completed")
	} else {
		log.Info("no need to init db data")
	}
	return nil

}
