package copyfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// получем имя и расширение (полное имя файла)
func (f *fileOriginal) getNameAndExp() (string, string) {
	arr := strings.Split(f.Name, ".")
	return arr[:len(arr)-1][0], arr[len(arr)-1]
}

// проверим одинакового ли размера эти файлы (true - если дубликат)
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

	// сравним размер файлов
	if stat.Size() == f.Size {
		return true, nil
	}

	return false, nil

}

// проверим есть ли файл с таким именем по заданному путь
func (f *fileOriginal) checkFileDoubleName() bool {
	// если файла с таким именем нет
	if _, err := os.Stat(filepath.Join(f.PathNew, f.NameNew)); os.IsNotExist(err) {
		return false
	}

	return true
}

// проверка файла на уникальность и замена имени при совпадении
// возвращает true если можно копировать, false - если копировать не надо это дубликат
func (f *fileOriginal) createNewFileName() (bool, error) {

	// разделим имя и расширение
	name, exp := f.getNameAndExp()
	// переменная для нового имени
	var newName = f.Name

	// счетчик количества провери имен
	c := 0

	// добавочная к новому имени файла
	i := 2

	for {
		c++
		// число попыток создать новое имя
		if i > 100 {
			return false, fmt.Errorf("Использованно Максимально количество попыток переиминовать файл: %s ", filepath.Join(f.PathNew, f.Name))
		}
		// если это не перва итерация,
		// значит родное мя не подошло,
		// подбираем новое имя
		if c != 1 {
			// создаем новое имя
			newName = fmt.Sprintf("%s_%d.%s", name, i, exp)
			// сохраняем новое имя
			f.NameNew = newName
			i++
		}

		// если такого файла с именем нет, можно копировать
		if !f.checkFileDoubleName() {
			return true, nil
		}

		// если имя совпало проверяем дубликат ли это
		d, err := f.comparsion()
		// если произошла ошибка во время дубликата просто создаем новое имя
		if err != nil {
			continue
		}

		// если это дубликат
		if d {
			return false, nil
		}

	}

}
