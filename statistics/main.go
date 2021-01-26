package statistics

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
	"path/filepath"
	"photo_archiv/checktype"
)

// один файл
type FileOne struct {
	Path       string
	Name       string
	TypeString string
	TypeInt    int
	Size       float64
}

// список всех файлов
var FileList []FileOne

// размер файла в Mb
func sizeMb(s int64) float64 {

	var y float64 = float64(s) / 1024 / 1024
	return math.Round(y*100) / 100

}

// получаем данные о файле
func fileInfo(p string) {

	// получение инфо файл с открытием (дорого)
	//original, err := os.Open(p)
	//if err != nil {
	//	log.Fatalf("open: %v", err)
	//}
	//defer original.Close()
	//
	//stat, err := original.Stat()
	//if err != nil {
	//	log.Fatalf("stat: %v", err)
	//}

	// получение инфо файл без открытия
	stat, err := os.Lstat(p)
	if err != nil {
		log.Fatalf("stat: %v", err)
	}

	// получем тип в строковом виде и int const
	it, st := checktype.GetType(stat.Name())

	fi := FileOne{
		Path:       p,
		Name:       stat.Name(),
		TypeString: st,
		TypeInt:    it,
		Size:       sizeMb(stat.Size()),
	}

	// добавляем в список
	FileList = append(FileList, fi)

}

// общая статистика
type StatisticsFull struct {
	GeneralSize  float64
	FileCount    int64
	VideoCount   int64
	VideoSize    float64
	ImgCount     int64
	ImgSize      float64
	UnknownCount int64
	UnknownSize  float64
}

// размер директории
func DirSize(path string) (int64, error) {

	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// анализ всех найденных файлов в слайсе
func GetStatistic(fls []FileOne, p string) {

	st := StatisticsFull{}
	sd, _ := DirSize(p)

	for _, v := range fls {

		st.GeneralSize = st.GeneralSize + v.Size
		st.FileCount++

		switch v.TypeInt {
		case checktype.VIDEO:
			st.VideoCount++
			st.VideoSize = st.VideoSize + v.Size
		case checktype.PHOTO:
			st.ImgCount++
			st.ImgSize = st.ImgSize + v.Size
		case checktype.UNKNOWN:
			st.UnknownCount++
			st.UnknownSize = st.UnknownSize + v.Size
		}

	}
	fmt.Printf("\nРазмер директории: %.2f", sizeMb(sd))
	fmt.Printf(`
Всего файлов: %d        ошбий размер:  %.2f Мб 
Видео файлов: %d        общий размер:  %.2f Мб 
Фото: %d                общий размер:  %.2f Мб
Неизвестных файлов: %d  общий размер:  %.2f Мб 

`,

		st.FileCount, st.GeneralSize,
		st.VideoCount, st.VideoSize,
		st.ImgCount, st.ImgSize,
		st.UnknownCount, st.UnknownSize)

	fmt.Println("Анализ завершен!!!")

}

// перебираем файлы
func GetFileList(d string) []FileOne {

	// читаем путь
	files, err := ioutil.ReadDir(d)
	if err != nil {
		log.Printf("Путь не прочитан: %v", err)
	}

	for _, file := range files {

		// обновляем путь
		n := path.Join(d, file.Name())

		// если каталог
		if file.IsDir() {

			GetFileList(n)
			continue
		}

		// обробатываем файл
		fileInfo(n)
	}

	return FileList
}
