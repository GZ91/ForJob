package service

import (
	"context"
	"github.com/GZ91/linkreduct/internal/errorsapp"
	"github.com/GZ91/linkreduct/internal/models"
	"regexp"
)

func New(ctx context.Context, db Storeger, conf ConfigerService, ChsURLForDel chan []models.StructDelURLs) *NodeService {
	Node := &NodeService{
		db:           db,
		conf:         conf,
		URLFormat:    regexp.MustCompile(`^(?:https?:\/\/)`),
		URLFilter:    regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?(\w+\.[^:\/\n]+)`),
		ChsURLForDel: ChsURLForDel,
	}

	return Node
}

type NodeService struct {
	db           Storeger
	conf         ConfigerService
	URLFormat    *regexp.Regexp
	URLFilter    *regexp.Regexp
	ChsURLForDel chan []models.StructDelURLs
}

func (r *NodeService) addURL(ctx context.Context, link string) (string, error) {
	return r.db.AddURL(ctx, link)
}

func (r *NodeService) getFormatLongLink(longLink string) (string, error) {
	if !r.URLFormat.MatchString(longLink) {
		longLink = "http://" + longLink
	}
	return longLink, nil
}

func (r *NodeService) Ping(ctx context.Context) error {
	return r.db.Ping(ctx)
}

func (r *NodeService) GetURLsToken(ctx context.Context, token string) ([]models.ReturnedStructURL, error) {
	addressServer := r.conf.GetAddressServerURL()
	returnedStructURL, err := r.db.GetLinksToken(ctx, token)
	if err != nil {
		return nil, err
	}
	for index, val := range returnedStructURL {
		returnedStructURL[index].ShortURL = addressServer + val.ShortURL
	}
	return returnedStructURL, nil
}

func (r *NodeService) Close() error {
	close(r.ChsURLForDel)
	return nil
}

func (r *NodeService) AddBatchLink(ctx context.Context, batchLink []string) (map[string]string, error) {

	for _, data := range batchLink {
		if !r.URLFilter.MatchString(data) {
			return nil, errorsapp.ErrInvalidLinkReceived
		}
	}

	return r.db.AddBatchLink(ctx, batchLink)
}
