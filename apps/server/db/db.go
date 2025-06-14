package db

import (
	"database/sql"
	_ "embed"

	"go.uber.org/zap"
)

//go:embed init.sql
var initSql string

//go:embed model.sql
var modelSql string

func InitDb(db *sql.DB, log *zap.Logger) error {
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
