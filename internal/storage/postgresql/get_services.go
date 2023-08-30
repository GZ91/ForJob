package postgresql

import "context"

func (Node *DB) GetServices(ctx context.Context, name string) (map[string]string, error) {
	data := make(map[string]string)
	if name == "" {
		rows, err := Node.db.Query("SELECT token, nameservice FROM tokens")
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var token, nameservice string
			rows.Scan(&token, &nameservice)
			data[nameservice] = token
		}
		return data, nil
	} else {
		rows, err := Node.db.Query("SELECT token, nameservice FROM  tokens WHERE nameservice = $1", name)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var token, nameservice string
			rows.Scan(&token, &nameservice)
			data[nameservice] = token
		}
		return data, nil
	}
	return nil, nil
}
