package checktype

import "strings"

// получем имя и расширение (полное имя файла)
func parseNameFile(n string) (string, string) {
	arr := strings.Split(n, ".")
	return arr[:len(arr)-1][0], arr[len(arr)-1]
}

// получаем тип
func check(e string) int {
	for _, f := range formatPhoto {
		if e == f {
			return PHOTO
		}
	}

	for _, f := range formatVideo {
		if e == f {
			return VIDEO
		}
	}

	return UNKNOWN
}

func GetType(fileName string) (int, string) {

	_, exp := parseNameFile(fileName)

	exp = strings.ToLower(exp)

	switch check(exp) {
	case VIDEO:
		return VIDEO, "видео"
	case PHOTO:
		return PHOTO, "фото"
	default:
		return UNKNOWN, "не определенно"

	}

}
