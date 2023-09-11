package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"go.uber.org/zap"
)

func (d *DB) FindLongURL(ctx context.Context, OriginalURL string, token string) (string, bool, error) {
	con, err := d.db.Conn(ctx)
	if err != nil {
		logger.Log.Error("failed to connect to the database", zap.Error(err))
		return "", false, err
	}
	defer con.Close()
	row := con.QueryRowContext(ctx, `SELECT ShortURL
	FROM short_origin_reference WHERE OriginalURL = $1 AND token = $2 limit 1`, OriginalURL, token)
	var shortURL string
	err = row.Scan(&shortURL)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log.Error("when scanning the result of the original link search query", zap.Error(err))
		return "", false, err
	}
	if shortURL != "" {
		return shortURL, true, nil
	}
	return "", false, nil
}
