package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	ErrBadLink = errors.New("bad link")
)

func main() {
	flag.Parse()
	link := flag.Arg(0)

	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err.Error())
	}

	filename, err := makeDirName(link)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = os.Mkdir(filename, os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	f, err := os.Create(filename + "/index.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	f.WriteString(string(bytes))
}

// makeDirName задает имя папки, куда будет сохранен файл сайта
func makeDirName(link string) (string, error) {
	args := strings.Split(link, "//")
	if len(args) < 2 {
		return "", ErrBadLink
	}

	return strings.Split(args[1], "/")[0], nil
}
