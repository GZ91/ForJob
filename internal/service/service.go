package service

import (
<<<<<<< HEAD
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
=======
	"encoding/json"
	"github.com/GZ91/EmploymentTest/internal/app/logger"
	"go.uber.org/zap"
)

type Storage interface {
	GetColumn() []string
	FindLines(string) ([][]string, error)
}

type Service struct {
	NodeStorage Storage
}

func New(storage Storage) (*Service, error) {
	return &Service{
		NodeStorage: storage,
	}, nil
}

func (s Service) RevertSearchStructures(id string) ([][]byte, error) {

	data, err := s.NodeStorage.FindLines(id)
	if err != nil {
		logger.Log.Error("При попытке найти строки по идентификатору произошла ошибка", zap.Error(err))
	}
	countData := len(data)
	returnMap := make([](map[string]string), 0)
	for i := 0; i < countData; i++ {
		returnMap = append(returnMap, make(map[string]string))
	}

	columns := s.NodeStorage.GetColumn()
	var dataReturn [][]byte
	for line := 0; line < countData; line++ {
		for column := 0; column < len(columns); column++ {
			returnMap[line][columns[column]] = data[line][column]
		}
		data, err := json.Marshal(returnMap)
		if err != nil {
			logger.Log.Error("при попытке конвертирования произошла ошибка", zap.Error(err))
			return nil, err
		}
		dataReturn = append(dataReturn, data)
	}

	return dataReturn, nil
>>>>>>> master
}
