package main

import (
	"encoding/json"
	"log"
	"os"
)

type SettingF struct {
	HomeDir   string `json:"home_dir"`
	NewDir    string `json:"new_dir"`
	VideoCopy bool   `json:"video_copy"`
}

// читаем настройки
func setting() *SettingF {

	sFile, err := os.Open("setting.json")
	if err != nil {
		log.Fatal(err)
	}
	defer sFile.Close()

	sFileDecoder := json.NewDecoder(sFile)
	if err != nil {
		log.Fatal(err)
	}

	s := &SettingF{}

	err = sFileDecoder.Decode(&s)
	if err != nil {
		log.Fatal(err)
	}

	return s

}
