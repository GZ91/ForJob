package service

import (
	"context"
)

func (s *NodeService) GetNewTokens(ctx context.Context, NameServices []string) (map[string]string, error) {
	return s.db.GetTokens(ctx, NameServices)
}
