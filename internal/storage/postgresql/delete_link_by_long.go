package postgresql

import (
	"context"
	"github.com/GZ91/linkreduct/internal/errorsapp"
)

func (d *DB) DeleteLinkByLongLink(ctx context.Context, longLink string, token string) error {
	tx, err := d.db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	tx.Exec("SELECT * FROM short_origin_reference WHERE OriginalURL = $1 AND token = $2 FOR UPDATE", longLink, token)
	row := tx.QueryRow("SELECT COUNT(*) FROM short_origin_reference WHERE OriginalURL = $1 AND token = $2", longLink, token)
	var countToken int
	row.Scan(&countToken)
	if countToken == 0 {
		return errorsapp.ErrNotFoundLink
	}
	tx.Exec("DELETE FROM short_origin_reference WHERE OriginalURL = $1 AND token = $2", longLink, token)
	tx.Commit()
	return nil
}
