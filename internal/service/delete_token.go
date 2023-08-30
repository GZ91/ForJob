package service

import "context"

func (r *NodeService) DeleteToken(ctx context.Context, token string) error {
	return r.db.DeleteToken(ctx, token)
}
