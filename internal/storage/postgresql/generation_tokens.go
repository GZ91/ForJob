package postgresql

import (
	"context"
	"github.com/google/uuid"
)

func (db *DB) GetTokens(ctx context.Context, namesServices []string) (map[string]string, error) {
	tx, err := db.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	returnData := make(map[string]string)
	for _, nameService := range namesServices {
		_, err := tx.Exec("SELECT * FROM tokens WHERE nameservice = $1 FOR UPDATE", nameService)
		if err != nil {
			return nil, err
		}
		row := tx.QueryRow("SELECT COUNT(nameservice) FROM tokens WHERE nameservice = $1", nameService)
		var countService int
		row.Scan(&countService)
		if countService == 0 {
			token := uuid.NewString()
			_, err := tx.Exec("INSERT INTO tokens (nameservice, token) VALUES ($1, $2)", nameService, token)
			if err != nil {
				return nil, err
			}
			returnData[nameService] = token
		} else {
			row := tx.QueryRow("SELECT service, token FROM tokens WHERE nameservice = $1 LIMIT 1", nameService)
			var service, token string
			row.Scan(&service, &token)
			returnData[service] = token
		}
	}
	tx.Commit()
	return returnData, nil
}
