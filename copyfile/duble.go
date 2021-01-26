package copyfile

import (
	"fmt"
	"os"
	"path/filepath"
)

// сравнение размеров
func (f *fileOriginal) comparsion() (bool, error) {

	original, err := os.Open(filepath.Join(f.PathNew, f.NameNew))
	if err != nil {
		return false, fmt.Errorf("Проверка на дубликат, не удалось открыть файл: %v", err)
	}
	defer original.Close()

	stat, err := original.Stat()
	if err != nil {
		return false, fmt.Errorf("Проверка на дубликат, не удалось получить данные файла: %v", err)
	}

	if stat.Size() == f.Size {
		return true, nil
	}

	return false, nil

}

// проверка есть ли такой файл с таким именем
func (f *fileOriginal) checkDuble() error {

	// если файла с таким именем нет
	if _, err := os.Stat(filepath.Join(f.PathNew, f.NameNew)); os.IsNotExist(err) {
		return nil
	}

	// если есть проверим может это тот же файл
	duplicate, err := f.comparsion()
	if err != nil {
		return err
	}

	// если дубликат
	if duplicate {
		return nil
	}

	return nil
}
