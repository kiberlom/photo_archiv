package copyfile

import (
	"fmt"
	"photo_archiv/checktype"
)

func (f *fileOriginal) checkSetting() error {

	// проверка на тип
	if !settingOsn.VideoCopy && f.Type == checktype.VIDEO {
		return fmt.Errorf(fmt.Sprintf("%s Запрещено копировать ВИДЕО", f.Path))
	}

	if !settingOsn.ImgCopy && f.Type == checktype.PHOTO {
		return fmt.Errorf(fmt.Sprintf("%s Запрещено копировать ИЗОБРАЖЕНИЕ", f.Path))
	}

	if !settingOsn.UnknownCopy && f.Type == checktype.UNKNOWN {
		return fmt.Errorf(fmt.Sprintf("%s Запрещено копировать НЕИЗВЕСТНЫЕ", f.Path))
	}

	return nil

}
