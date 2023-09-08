package service

import (
	"context"
)

func (r *NodeService) DeleteLinkByLongLink(ctx context.Context, longLink string, token string) error {
	return r.db.DeleteLinkByLongLink(ctx, longLink, token)
}
