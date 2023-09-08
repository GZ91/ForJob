package service

import (
	"context"
)

func (Node *NodeService) GetLinks(ctx context.Context, token string) (map[string]string, error) {
	data, err := Node.db.GetLinks(ctx, token)
	if err != nil {
		return nil, err
	}
	returnMap := make(map[string]string)
	for short, long := range data {
		returnMap[Node.conf.GetAddressServerURL()+short] = long
	}
	return returnMap, nil
}
