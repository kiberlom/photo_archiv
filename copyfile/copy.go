package copyfile

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func (f *fileOriginal) copyFile() error {

	original, err := os.Open(f.Path)
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	defer original.Close()

	new, err := os.Create(filepath.Join(f.PathNew, f.NameNew))
	if err != nil {
		return err
	}
	defer new.Close()

	_, err = io.Copy(new, original)
	if err != nil {
		return err
	}

	return nil

}

func (f *fileOriginal) copyFileProgress() error {

	return nil

}

func countByte(f string) int {
	input, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return len(input)
}

func uuu() {

	f := "E:\\001.mp4"
	c := "E:\\001_copy.mp4"

	b := countByte(f)

	source, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()

	destination, err := os.Create(c)
	if err != nil {
		log.Fatal(err)
	}
	defer destination.Close()

	buf := make([]byte, b/9)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			log.Fatal(err)
		}

		fmt.Print("=")
	}

	fmt.Print(">")

}
