package service

import (
	"context"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/errorsapp"
	"github.com/GZ91/linkreduct/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"regexp"
)

// Storeger
//
//go:generate mockery --name Storeger --with-expecter
type Storeger interface {
	AddURL(context.Context, string) (string, error)
	GetURL(context.Context, string) (string, bool, error)
	Ping(context.Context) error
	AddBatchLink(context.Context, []models.IncomingBatchURL) ([]models.ReleasedBatchURL, error)
	FindLongURL(context.Context, string) (string, bool, error)
	GetLinksToken(context.Context, string) ([]models.ReturnedStructURL, error)
	InitializingRemovalChannel(context.Context, chan []models.StructDelURLs) error
}

// Storeger
//
//go:generate mockery --name ConfigerService --with-expecter
type ConfigerService interface {
	GetAddressServerURL() string
	GetSecretKey() string
}

func New(ctx context.Context, db Storeger, conf ConfigerService, ChsURLForDel chan []models.StructDelURLs) *NodeService {
	Node := &NodeService{
		db:           db,
		conf:         conf,
		URLFormat:    regexp.MustCompile(`^(?:https?:\/\/)`),
		URLFilter:    regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?(\w+\.[^:\/\n]+)`),
		ChsURLForDel: ChsURLForDel,
	}

	err := Node.db.InitializingRemovalChannel(ctx, Node.ChsURLForDel)
	if err != nil {
		logger.Log.Error("error when initializing the delete link channel", zap.Error(err))
		return nil
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

func (r *NodeService) GetURL(ctx context.Context, id string) (string, bool, error) {
	return r.db.GetURL(ctx, id)
}

func (r *NodeService) addURL(ctx context.Context, link string) (string, error) {
	return r.db.AddURL(ctx, link)
}

func (r *NodeService) GetSmallLink(ctx context.Context, longLink string) (string, error) {

	longLink, err := r.getFormatLongLink(longLink)
	if err != nil {
		return "", err
	}
	id, ok, err := r.db.FindLongURL(ctx, longLink)
	if err != nil {
		return "", err
	}
	if ok {
		return r.conf.GetAddressServerURL() + id, errorsapp.ErrLinkAlreadyExists
	}
	id, err = r.addURL(ctx, longLink)
	if err != nil {
		return "", err
	}
	return r.conf.GetAddressServerURL() + id, nil
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

func (r *NodeService) DeletedLinks(listURLs []string, userID string) {

	var dataForDel []models.StructDelURLs
	for _, val := range listURLs {
		data := models.StructDelURLs{URL: val, UserID: userID}
		dataForDel = append(dataForDel, data)
	}

	r.ChsURLForDel <- dataForDel
}

func (r *NodeService) Close() error {
	close(r.ChsURLForDel)
	return nil
}

func (r *NodeService) GetDataToken() (string, string, error) {
	return getDataForToken(r.conf.GetSecretKey())
}

func getDataForToken(secretKey string) (string, string, error) {
	Token := uuid.New().String()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		Token: Token,
	})

	DataString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", "", err
	}

	return DataString, Token, nil
}
