package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type cmpFile struct {
	pathNew string
	nameNew string
	original os.FileInfo
}

// сравнение размеров
func (c *cmpFile) comparsion() (bool,error) {

	original, err := os.Open(filepath.Join(c.pathNew, c.nameNew))
	if err != nil {
		return false, fmt.Errorf("Проверка на дубликат, не удалось открыть файл: %v", err)
	}
	defer original.Close()

	stat, err := original.Stat()
	if err != nil {
		return false, fmt.Errorf("Проверка на дубликат, не удалось получить данные файла: %v", err)
	}


	if stat.Size() == c.original.Size(){
		return true, nil
	}

	return false, nil


}