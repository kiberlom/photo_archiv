package statistics

import (
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
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

	original, err := os.Open(p)
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	defer original.Close()

	stat, err := original.Stat()
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
	GeneralSize float64
	FileCount   int64
	VideoCount  int64
	VideoSize   float64
	ImgCount    int64
	ImgSize     float64
}

func getStatistic() {

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
