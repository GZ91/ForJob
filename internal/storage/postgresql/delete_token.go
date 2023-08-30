package postgresql

import (
	"context"
	"github.com/GZ91/linkreduct/internal/errorsapp"
)

func (d *DB) DeleteToken(ctx context.Context, token string) error {
	tx, err := d.db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	tx.Exec("SELECT * FROM tokens WHERE token = $1 FOR UPDATE", token)
	row := tx.QueryRow("SELECT COUNT(*) FROM tokens WHERE token = $1", token)
	var countToken int
	row.Scan(&countToken)
	if countToken == 0 {
		return errorsapp.ErrNotFoundToken
	}
	tx.Exec("DELETE FROM tokens WHERE token = $1", token)
	tx.Commit()
	return nil
}
