package service

import (
	"context"
)

func (s *NodeService) GetTokens(ctx context.Context, NameServices []string) (map[string]string, error) {
	return s.db.GetTokens(ctx, NameServices)
}
