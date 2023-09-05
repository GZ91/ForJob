package service

import (
	"context"
)

func (Node *NodeService) GetLinks(ctx context.Context, token string) (map[string]string, error) {
	return Node.db.GetLinks(ctx, token)
}
