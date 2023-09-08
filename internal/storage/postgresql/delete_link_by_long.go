package postgresql

import (
	"context"
	"github.com/GZ91/linkreduct/internal/errorsapp"
)

func (d *DB) DeleteLinkByLongLink(ctx context.Context, longLink string) error {
	tx, err := d.db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	tx.Exec("SELECT * FROM short_origin_reference WHERE OriginalURL = $1 FOR UPDATE", longLink)
	row := tx.QueryRow("SELECT COUNT(*) FROM short_origin_reference WHERE OriginalURL = $1", longLink)
	var countToken int
	row.Scan(&countToken)
	if countToken == 0 {
		return errorsapp.ErrNotFoundLink
	}
	tx.Exec("DELETE FROM short_origin_reference WHERE OriginalURL = $1", longLink)
	tx.Commit()
	return nil
}
