package main

import (
	"fmt"
	"os"
	"photo_archiv/copyfile"
	"photo_archiv/setting"
	"photo_archiv/statistics"
)

// настройки
var SettingRoot *setting.SettingF

// перебираем файлы
func osn(fl []statistics.FileOne) {

	// считаем проценты исходя из количества обработанных файлов
	ol := float64(len(fl))
	nl := float64(0)

	for _, f := range fl {

		copyfile.Start(f, SettingRoot)
		nl++

		// высчитываем проценты
		g := nl * 100 / ol

		fmt.Printf("%.2f%%  ", g)

	}

}

// чтобы посмотрать статистику сначала
func startCopy() {

	fmt.Print("Введите (Y) чтобы начать копировать: ")

	var a string
	fmt.Scan(&a)

	if a != "Y" {
		fmt.Println("Работа завершена ")
		os.Exit(0)

	}

}

func main() {

	// настройки
	SettingRoot = setting.GetSetting()

	fmt.Println("Подождите идет анализ...")
	// получаем список найденных файлов
	f := statistics.GetFileList(SettingRoot.HomeDir)
	// получаем общую статистику
	statistics.GetStatistic(f, SettingRoot.HomeDir)

	startCopy()

	osn(f)

}
