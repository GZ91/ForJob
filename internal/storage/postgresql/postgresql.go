package postgresql

import (
	"context"
	"database/sql"
	"github.com/GZ91/linkreduct/internal/app/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

type DB struct {
	conf           ConfigerStorage
	generatorRunes GeneratorRunes
	ps             string
	db             *sql.DB
}

func New(ctx context.Context, config ConfigerStorage, generatorRunes GeneratorRunes) (*DB, error) {
	db := &DB{conf: config, generatorRunes: generatorRunes}
	ConfDB := db.conf.GetConfDB()
	db.ps = ConfDB.StringServer
	err := db.openDB()
	if err != nil {
		return nil, err
	}
	err = db.createTables(ctx)
	if err != nil {
		return nil, err
	}
	return db, err
}

func (d *DB) openDB() error {
	db, err := sql.Open("pgx", d.ps)
	if err != nil {
		logger.Log.Error("failed to open the database", zap.Error(err))
		return err
	}
	d.db = db
	return nil
}

func (d *DB) Ping(ctx context.Context) error {
	con, err := d.db.Conn(ctx)
	if err != nil {
		logger.Log.Error("failed to connect to the database", zap.Error(err))
		return err
	}
	defer con.Close()
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
