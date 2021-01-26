package copyfile

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"os"
	"path/filepath"
	"photo_archiv/checktype"
	"photo_archiv/setting"
	"photo_archiv/statistics"
	"strconv"
)

// файл
type fileOriginal struct {
	Path               string
	Name               string
	Size               int64
	Type               int
	TypeStr            string
	YearCreate         int
	MounthCreate       int
	MounthCreateRusStr string
	DayCreate          int
	Open               *os.File
	PathNew            string
	NameNew            string
}

// настройки
var settingOsn *setting.SettingF

// получение дату exif (тоесть реальная дата создания )
func (f *fileOriginal) getDateExif() error {

	// получем данные exif
	x, err := exif.Decode(f.Open)
	if err != nil {
		return fmt.Errorf("Невозможно получить данные exif, файл [%s]: %v ", f.Path, err)
	}

	// получаем дату
	d, err := x.DateTime()
	if err != nil {
		return fmt.Errorf("Невозможно получить дату из exif, файл [%s]: %v ", f.Path, err)
	}

	// добавляем данные
	f.YearCreate = d.Year()
	f.MounthCreate = int(d.Month())
	f.DayCreate = d.Day()

	return nil

}

// получение дату exif (тоесть реальная дата создания )
func (f *fileOriginal) getDateStat() error {

	stat, err := f.Open.Stat()
	if err != nil {
		return fmt.Errorf("Невозможно получить простую дату, файл [%s]: %v ", f.Path, err)
	}

	// добавляем данные
	f.YearCreate = stat.ModTime().Year()
	f.MounthCreate = int(stat.ModTime().Month())
	f.DayCreate = stat.ModTime().Day()

	return nil

}

// открываем файл
func (f *fileOriginal) openFile() error {

	// открываем файл
	original, err := os.Open(f.Path)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Невозможно открыть файл [%s]: ", f.Path), err)
	}

	// сохраняем открытие
	f.Open = original

	return nil
}

// получим размер в байтах
func (f *fileOriginal) getSize() error {
	stat, err := f.Open.Stat()
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Невозможно получить размер файла [%s]: ", f.Path), err)
	}

	f.Size = stat.Size()
	return nil
}

// получим месяц на русском
func (f *fileOriginal) mounthRusStr() {

	for i, r := range mounth {
		if i == f.MounthCreate {
			f.MounthCreateRusStr = r
			return
		}
	}

	f.MounthCreateRusStr = "не определенн"
}

// определяем тип
func (f *fileOriginal) getType() {

	i, s := checktype.GetType(f.Name)
	f.Type = i
	f.TypeStr = s

}

// создание каталога (новый путь к файлу)
func createNewDir(p string) error {

	err := os.MkdirAll(p, 0)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

// создаем каталог по новуму пуьт
func (f *fileOriginal) newPath() error {
	// новый путь
	p := filepath.Join(settingOsn.NewDir, strconv.Itoa(f.YearCreate), f.MounthCreateRusStr, f.TypeStr)
	// записываем новый путь
	f.PathNew = p

	// если каталог уже есть
	if _, err := os.Stat(p); !os.IsNotExist(err) {
		return nil
	}

	// создаем каталог
	if err := createNewDir(p); err != nil {
		return fmt.Errorf(fmt.Sprintf("Невозможно создать каталог [%s]: ", p), err)
	}

	return nil
}

// НАЧАЛО
func Start(fl statistics.FileOne, s *setting.SettingF) {

	// получем настройки
	settingOsn = s

	// создаем структуру файла
	f := fileOriginal{
		Path:               fl.Path,
		Name:               fl.Name,
		Size:               0,
		Type:               0,
		TypeStr:            "",
		YearCreate:         0,
		MounthCreate:       0,
		MounthCreateRusStr: "",
		DayCreate:          0,
		Open:               nil,
		PathNew:            "",
		NameNew:            fl.Name,
	}

	// получем тип
	f.getType()

	// открываем файл
	if err := f.openFile(); err != nil {
		fmt.Println(err)
		return
	}
	defer f.Open.Close()

	// получим размер
	if err := f.getSize(); err != nil {
		fmt.Println(err)
		return
	}

	// получаем дату создания файла exif
	if err := f.getDateExif(); err != nil {
		// пытаемся получить простую дату
		if err := f.getDateStat(); err != nil {
			fmt.Println(err)
			return
		}
	}

	//получем месяц на русском
	f.mounthRusStr()

	// проверка настроек
	if err := f.checkSetting(); err != nil {
		fmt.Println(err)
		return
	}

	// проверяем и создаем каталог под файл
	if err := f.newPath(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", f)

}
