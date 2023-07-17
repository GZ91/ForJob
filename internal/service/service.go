package service

import (
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

func (s Service) RevertSearchStructures(id string) ([]byte, error) {

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

	for line := 0; line < countData; line++ {
		for column := 0; column < len(columns); column++ {
			returnMap[line][columns[column]] = data[line][column]
		}
	}
	return json.Marshal(returnMap)
}
