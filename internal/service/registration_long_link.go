package service

import (
	"context"
	"github.com/GZ91/linkreduct/internal/errorsapp"
)

func (r *NodeService) AddBatchLink(ctx context.Context, batchLink []string) (releasedBatchURL map[string]string, errs error) {

	for _, data := range batchLink {
		link := data.OriginalURL

		if !r.URLFilter.MatchString(link) {
			return nil, errorsapp.ErrInvalidLinkReceived
		}
	}

	releasedBatchURL, errs = r.db.AddBatchLink(ctx, batchLink)
	for index, val := range releasedBatchURL {
		releasedBatchURL[index].ShortURL = r.conf.GetAddressServerURL() + val.ShortURL
	}
	return
}
