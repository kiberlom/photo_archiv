package main

import (
	"fmt"
	"photo_archiv/copyfile"
	"photo_archiv/setting"
	"photo_archiv/statistics"
)

// настройки
var SettingRoot *setting.SettingF

// перебираем файлы
func osn(fl []statistics.FileOne) {

	ol := len(fl)
	nl := 0

	for _, f := range fl {

		copyfile.Start(f, SettingRoot)
		nl++

		g := float64(nl) * 100 / float64(ol)

		fmt.Printf("%.2f%%", g)

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

	osn(f)

}
