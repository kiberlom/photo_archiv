package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"photo_archiv/checktype"
	"photo_archiv/statistics"
)

// настройки
var settingRoot *SettingF

// перебираем файлы
func dirRande(d string) {

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

			dirRande(n)
			continue
		}

		// обробатываем файл
		fileTE(n)
	}
}

func analysisStat(fl []statistics.FileOne) {

	var su float64 = 0

	var video int64 = 0
	var videoSize float64 = 0

	var pf int64 = 0
	var pfSize float64 = 0

	var nf int64 = 0
	var nfSize float64 = 0

	for _, v := range fl {
		//fmt.Printf("name: %s  type: %s  size: %.2f  path:%s\n", v.Name, v.TypeString, v.Size, v.Path)
		su = su + v.Size

		switch v.TypeInt {
		case checktype.VIDEO:
			video++
			videoSize = videoSize + v.Size
		case checktype.PHOTO:
			pf++
			pfSize = pfSize + v.Size
		case checktype.NOT_FOUND:
			nf++
			nfSize = nfSize + v.Size
		}

	}

	fmt.Printf(`
Ошбий размер: %.2f Мб
Видео файлов: %d  общий размер:  %.2f Мб 
Фото: %d  общий размер:  %.2f Мб
Неизвестных файлов: %d  общий размер:  %.2f Мб 
`,

		su,
		video, videoSize,
		pf, pfSize,
		nf, nfSize)
	fmt.Println("Анализ завершен")

}

func main() {

	// настройки
	settingRoot = setting()

	// запускаем
	//dirRande(settingRoot.HomeDir)
	fmt.Println("Подождите идет анализ...")
	fl := statistics.GetFileList(settingRoot.HomeDir)

	analysisStat(fl)

}
