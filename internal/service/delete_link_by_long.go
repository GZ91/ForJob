package service

import (
	"context"
)

func (r *NodeService) DeleteLinkByLongLink(ctx context.Context, longLink string) error {
	return r.db.DeleteLinkByLongLink(ctx, longLink)
}
