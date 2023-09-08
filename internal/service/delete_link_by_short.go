package service

import (
	"context"
	"strings"
)

func (r *NodeService) DeleteLinkByShortLink(ctx context.Context, shortLink string) error {
	shortLink = strings.Replace(shortLink, r.conf.GetAddressServerURL(), "", -1)
	return r.db.DeleteLinkByShortLink(ctx, shortLink)
}
