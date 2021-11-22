package main

import (
	"fmt"
	"io"
	"os"
	"flag"
)

func main() {
	
	var path string
	var bytes int
    flag.StringVar(&path, "path", "log.txt", "Необходимо ввести путь к файлу, который нужно открыть")
	flag.IntVar(&bytes, "byte", 1024, "Необходимо ввести максимальное количество байт, которое будем считывать из файла за раз")
	
	/*
	bytes := flag.Int("bytes", 1024, "Необходимо ввести максимальное количество байт, которое будем считывать из файла за раз") --> buf := make([]byte, *bytes)
    path := flag.String("path", "/Users/aleksejvejcman/учеба/goland/log.txt", "Необходимо ввести путь к файлу ") --> os.Open(*path)
	*/
	
	flag.Parse()
	
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	buf := make([]byte, bytes)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}
