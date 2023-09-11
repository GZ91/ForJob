package service

import (
	"context"
	"regexp"
)

func New(ctx context.Context, db Storeger, conf ConfigerService) *NodeService {
	Node := &NodeService{
		db:        db,
		conf:      conf,
		URLFormat: regexp.MustCompile(`^(?:https?:\/\/)`),
		URLFilter: regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?(\w+\.[^:\/\n]+)`),
	}

	return Node
}

type NodeService struct {
	db        Storeger
	conf      ConfigerService
	URLFormat *regexp.Regexp
	URLFilter *regexp.Regexp
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

func (r *NodeService) Close() error {
	return nil
}
