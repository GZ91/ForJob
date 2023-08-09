package postgresql

import (
	"context"
	"database/sql"
)

func (d *DB) createTables(ctx context.Context) error {
	con, err := d.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer con.Close()
	err = createTableLinks(ctx, con)
	if err != nil {
		return err
	}
	err = createTableTokens(ctx, con)
	if err != nil {
		return err
	}

	return err
}

func createTableLinks(ctx context.Context, con *sql.Conn) error {
	_, err := con.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS short_origin_reference 
(
	id serial PRIMARY KEY,
	token VARCHAR(45)  NOT NULL,
	ShortURL VARCHAR(250) NOT NULL,
	OriginalURL TEXT
);`)
	return err
}

func createTableTokens(ctx context.Context, con *sql.Conn) error {
	_, err := con.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS tokens 
(
	id serial PRIMARY KEY,
	token VARCHAR(45)  NOT NULL,
	nameservice VARCHAR(1000) NOT NULL,
);`)
	return err
}
