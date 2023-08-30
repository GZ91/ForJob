package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/models"
	"go.uber.org/zap"
)

func (d *DB) AddURL(ctx context.Context, URL string) (string, error) {
	var token string
	var tokenIDCTX models.CtxString = "Authorization"
	UserIDVal := ctx.Value(tokenIDCTX)
	if UserIDVal != nil {
		token = UserIDVal.(string)
	}

	con, err := d.db.Conn(ctx)
	if err != nil {
		logger.Mserror("failed to connect to the database", err, nil)
		return "", err
	}
	defer con.Close()
	lenShort := d.conf.GetStartLenShortLink()
	index := 0

	var shorturl string
	for {
		shorturl = d.generatorRunes.RandStringRunes(lenShort)
		row := con.QueryRowContext(ctx, "SELECT COUNT(id) FROM short_origin_reference WHERE shorturl = $1", shorturl)
		var countShorturl int
		err := row.Scan(&countShorturl)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			logger.Log.Error("when scanning a request for a shortcut", zap.Error(err))
			return "", err
		}
		if countShorturl == 0 {
			break
		}
		index++
		if index == d.conf.GetMaxIterLen() {
			lenShort++
			index = 0
		}
	}

	_, err = con.ExecContext(ctx, "INSERT INTO short_origin_reference(token, shorturl, originalurl) VALUES ($1, $2, $3);",
		token, shorturl, URL)
	if err != nil {
		logger.Mserror("error when adding a record to the database", err, nil)
		return "", err
	}
	return shorturl, nil
}
