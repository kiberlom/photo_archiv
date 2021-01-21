package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// типы файлов
const (
	VIDEO     = "Видео"
	PHOTO     = "Фото"
	NOT_FOUND = "Не определенно"
)

type infoFile struct {
	mounth     int
	year       int
	name       string
	exp        string
	pathOrigin string
}

// проверка идентичности файла (новый путь, новое имя) (есть ли имя, дубликат, ошибка)
func checkOriginal(path, name string, org os.FileInfo) (bool, bool, error) {

	com := cmpFile{
		pathNew:  path,
		nameNew:  name,
		original: org,
	}

	// если файла с таким именем нет
	if _, err := os.Stat(filepath.Join(path, name)); os.IsNotExist(err) {
		return false, false, nil
	}

	// если есть проверим может это тот же файл
	duplicate, err := com.comparsion()
	if err != nil {
		return false, false, err
	}

	// если дубликат
	if duplicate {
		return true, true, nil
	}

	return true, false, nil
}

// копирование файла (новый путь новое имя, оригинал файла)
func copyFile(path, nameNew string, org *os.File) error {

	new, err := os.Create(filepath.Join(path, nameNew))
	if err != nil {
		return err
	}
	defer new.Close()

	_, err = io.Copy(new, org)
	if err != nil {
		return err
	}

	return nil

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

// месяц на русском (месяц - число)
func getMonthRus(m int) string {

	for i, r := range mounth {
		if i == m {
			return r
		}
	}

	return strconv.Itoa(m)

}

// получем имя и расширение (полное имя файла)
func parseNameFile(n string) (string, string) {
	arr := strings.Split(n, ".")
	return arr[:len(arr)-1][0], arr[len(arr)-1]
}

// определяем формат (расширение)
func formatFile(exp string) string {

	exp = strings.ToLower(exp)

	for _, f := range formatPhoto {
		if exp == f {
			return PHOTO
		}
	}

	for _, f := range formatVideo {
		if exp == f {
			return VIDEO
		}
	}

	return NOT_FOUND

}

//новый путь (годб месяц, формат)
func PathCreate(y, m int, f string) (p string, is bool) {

	// новый путь
	p = filepath.Join(settingRoot.NewDir, strconv.Itoa(y), getMonthRus(m), f)
	is = true

	// если каталог уже есть
	if _, err := os.Stat(p); !os.IsNotExist(err) {
		return
	}

	// создаем каталог
	if err := createNewDir(p); err != nil {
		is = false
		return
	}

	return

}

// инфо о файле
func fileTE(p string) {

	original, err := os.Open(p)
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	defer original.Close()

	stat, err := original.Stat()
	if err != nil {
		log.Fatalf("stat: %v", err)
	}

	// полное имя оригинала
	nameFile := stat.Name()

	// получаем имя и расширение
	nm, exp := parseNameFile(nameFile)

	// месяц
	m := int(stat.ModTime().Month())

	// год
	y := stat.ModTime().Year()

	// формат файла
	ff := formatFile(exp)

	// новый путь для файла
	pathNew, _ := PathCreate(y, m, ff)

	i := 2
	for {

		// проверим на совпадение имен и дубликаты
		nameTrue, dubl, err := checkOriginal(pathNew, nameFile, stat)

		if err != nil {
			fmt.Println("ФАЙЛ: ", stat.Name(), " ПУТЬ: ", filepath.Join(pathNew, nameFile), " ОШИБКА: ", err.Error())
			return
		}

		// если дубликат
		if dubl {
			fmt.Println("ФАЙЛ: ", stat.Name(), " ПУТЬ: ", filepath.Join(pathNew, nameFile), "   Дубликат")
			return
		}

		// если совпало только имя
		if nameTrue {
			nameFile = fmt.Sprintf("%s(%d).%s", nm, i, exp)
			i++
			continue
		}

		break
	}

	err = copyFile(pathNew, nameFile, original)
	if err != nil {
		s := fmt.Sprintf("ОШИБКА КОПИРОВАНИЯ: %s ====> %s", p, filepath.Join(pathNew, nameFile))
		fmt.Println(s)
		return
	}

	fmt.Println(fmt.Sprintf("%s ====> %s СКОПИРОВАН", p, filepath.Join(pathNew, nameFile)))

}
