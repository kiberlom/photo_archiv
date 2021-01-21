package main

import (
	"io/ioutil"
	"log"
	"path"
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

func main() {

	// настройки
	settingRoot = setting()

	// запускаем
	dirRande(settingRoot.HomeDir)

}
