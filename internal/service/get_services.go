package service

import "context"

func (Node *NodeService) GetServices(ctx context.Context, name string) (map[string]string, error) {
	return Node.db.GetServices(ctx, name)
}
