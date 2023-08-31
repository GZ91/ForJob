package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/GZ91/linkreduct/internal/app/logger"
)

func (d *DB) GetURL(ctx context.Context, shortURL string) (string, bool, error) {
	con, err := d.db.Conn(ctx)
	if err != nil {
		logger.Mserror("failed to connect to the database", err, nil)
		return "", false, err
	}
	defer con.Close()
	row := con.QueryRowContext(ctx, `SELECT originalurl 
	FROM short_origin_reference WHERE shorturl = $1 limit 1`, shortURL)
	var originurl string
	err = row.Scan(&originurl)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Mserror("when scanning the request for the original link", err, nil)
		return "", false, err
	}
	if originurl != "" {
		return originurl, true, nil
	}
	return "", false, nil
}
