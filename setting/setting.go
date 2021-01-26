package setting

import (
	"encoding/json"
	"log"
	"os"
)

type SettingF struct {
	HomeDir     string `json:"home_dir"`
	NewDir      string `json:"new_dir"`
	VideoCopy   bool   `json:"video_copy"`
	ImgCopy     bool   `json:"photo_copy"`
	UnknownCopy bool   `json:"unknown_copy"`
}

// читаем настройки
func GetSetting() *SettingF {

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
