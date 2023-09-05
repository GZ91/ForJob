package postgresql

import "context"

func (db *DB) GetLinks(ctx context.Context, token string) (map[string]string, error) {
	con, err := db.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer con.Close()
	rows, err := con.QueryContext(ctx, "SELECT ShortURL, OriginalURL FROM short_origin_reference WHERE token = $1", token)
	if err != nil {
		return nil, err
	}
	returnData := make(map[string]string)
	var shorURL, originURL string
	for rows.Next() {
		err = rows.Scan(&shorURL, &originURL)
		if err != nil {
			return nil, err
		}
		returnData[shorURL] = originURL
	}
	return returnData, nil
}
