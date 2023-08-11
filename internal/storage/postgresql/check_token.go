package postgresql

import "context"

func (r *DB) CheckToken(ctx context.Context, token string) (bool, error) {
	con, err := r.db.Conn(ctx)
	row := con.QueryRowContext(ctx, "SELECT count(*) FROM tokens WHERE token = $1", token)
	var countToken int
	err = row.Scan(&countToken)
	if err != nil {
		return false, err
	}
	if countToken == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
