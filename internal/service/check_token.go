package service

import "context"

func (node *NodeService) CheckToken(ctx context.Context, Token string) (bool, error) {
	return node.db.CheckToken(ctx, Token)
}
