package storage

import (
	"encoding/csv"
	"github.com/GZ91/EmploymentTest/internal/app/logger"
	"go.uber.org/zap"
	"os"
	"runtime"
	"sync"
)

type StorageCSV struct {
	path string
	csv  [][]string
}

func New(path string) (*StorageCSV, error) {
	file, err := os.Open(path)
	if err != nil {
		logger.Log.Error("При создании узла хранения", zap.Error(err))
		return nil, err
	}
	defer file.Close() // в первой итерации забыл закрыть файл - поторопился
	csvStorage, err := csv.NewReader(file).ReadAll()
	if err != nil {
		logger.Log.Error("При чтении csv файла", zap.Error(err))
		return nil, err
	}
	var Node StorageCSV
	Node.path = path
	Node.csv = csvStorage
	return &Node, nil
}

func (s StorageCSV) GetColumn() []string {
	var temp []string
	temp = append(temp, s.csv[0]...)
	return temp
}

func (s StorageCSV) getData() [][]string {
	var tempR [][]string
	temp2 := s.csv[1:]
	tempR = append(tempR, temp2...) // передаю копированием, т.к. get подразумевает не отдавать оригинал чтобы его не поменяли.
	return tempR
}

func (s StorageCSV) FindLines(id string) ([][]string, error) {

	ch := make(chan []string)
	exch := make(chan [][]string)
	data := s.getData()
	countLine := len(data)
	CountCPU := runtime.NumCPU()
	CountLineForCPU := int(countLine / CountCPU)
	var wg sync.WaitGroup
	go picker(ch, exch)
	for i := 0; i < CountLineForCPU; i++ {
		wg.Add(1)
		if i == CountCPU-1 {
			go findLinesOn(ch, id, data[CountLineForCPU*i:][:], &wg)
			break
		} else if i == 0 {
			go findLinesOn(ch, id, data[:CountLineForCPU][:], &wg)
			continue
		}
		go findLinesOn(ch, id, data[CountLineForCPU*i : CountLineForCPU*i+CountLineForCPU][:], &wg)
	}
	wg.Wait()
	dataReturn := <-exch
	close(exch)
	close(ch)
	return dataReturn, nil

}

func findLinesOn(ch chan<- []string, id string, data [][]string, wg *sync.WaitGroup) {
	for _, val := range data {
		if val[1] == id {
			ch <- val
		}
	}
	wg.Done()
}

func picker(ch <-chan []string, exch chan [][]string) {
	var data [][]string
	for {
		select {
		case exch <- data:
			return
		case c := <-ch:
			data = append(data, c)
		}
	}
}
