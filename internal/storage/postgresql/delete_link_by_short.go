package postgresql

import (
	"context"
	"github.com/GZ91/linkreduct/internal/errorsapp"
)

func (d *DB) DeleteLinkByShortLink(ctx context.Context, shortLink string, token string) error {
	tx, err := d.db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	tx.Exec("SELECT * FROM short_origin_reference WHERE ShortURL = $1 AND token = $2 FOR UPDATE", shortLink, token)
	row := tx.QueryRow("SELECT COUNT(*) FROM short_origin_reference WHERE ShortURL = $1 AND token = $2", shortLink, token)
	var countToken int
	row.Scan(&countToken)
	if countToken == 0 {
		return errorsapp.ErrNotFoundLink
	}
	tx.Exec("DELETE FROM short_origin_reference WHERE ShortURL = $1 AND token = $2", shortLink, token)
	tx.Commit()
	return nil
}
