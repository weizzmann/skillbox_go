package main

import (
	"fmt"
	//"io"
	"os"
)

func main() {
	file, err := os.Open("new.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	buf := make([]byte, fileStat.Size())
	//if _, err := io.ReadFull(file, buf); err != nil {
	if _, err := file.Read(buf); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", buf)
}