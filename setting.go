package main

import (
	"encoding/json"
	"log"
	"os"
)

type SettingF struct {
	HomeDir string `json:"home_dir"`
	NewDir  string `json:"new_dir"`
}

// читаем настройки
func setting() *SettingF{

	// open the file pointer
	sFile, err := os.Open("setting.json")
	if err != nil {
		log.Fatal(err)
	}
	defer sFile.Close()

	// create a new decoder
	sFileDecoder := json.NewDecoder(sFile)
	if err != nil {
		log.Fatal(err)
	}

	// initialize the storage for the decoded data
	s := &SettingF{}

	// decode the data
	err = sFileDecoder.Decode(&s)
	if err != nil {
		log.Fatal(err)
	}

	return s

}