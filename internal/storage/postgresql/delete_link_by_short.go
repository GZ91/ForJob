package postgresql

import (
	"context"
	"github.com/GZ91/linkreduct/internal/errorsapp"
)

func (d *DB) DeleteLinkByShortLink(ctx context.Context, shortLink string) error {
	tx, err := d.db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	tx.Exec("SELECT * FROM short_origin_reference WHERE ShortURL = $1 FOR UPDATE", shortLink)
	row := tx.QueryRow("SELECT COUNT(*) FROM short_origin_reference WHERE ShortURL = $1", shortLink)
	var countToken int
	row.Scan(&countToken)
	if countToken == 0 {
		return errorsapp.ErrNotFoundLink
	}
	tx.Exec("DELETE FROM short_origin_reference WHERE ShortURL = $1", shortLink)
	tx.Commit()
	return nil
}
