package main

import (
	"flag"
	"fmt"
	"github.com/GZ91/EmploymentTest/internal/api/server"
	"github.com/GZ91/EmploymentTest/internal/app/logger"
)

func main() {
	logger.Initializing("info")
	pathS := flag.String("p", "", "csv file address")
	flag.Parse()
	path := *pathS
	if path == "" {
		logger.Log.Error("Не указано имя файла.")
		fmt.Println("Укажите пожалуйста имя файла CSV")
		return
	}
	err := server.Run(path)
	if err != nil {
		panic(err) // "никогда не паникуйте", но я думаю что паникавать в функции main имеет смысл
	}
}
