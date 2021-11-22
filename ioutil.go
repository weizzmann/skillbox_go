package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main(){
	
	var b bytes.Buffer

	text := ""
	i := 1
	
	for {
		timeNow := time.Now().Format("2006-01-02 15:04:05") // 🤯 
		
		fmt.Println("Введите строку (для выхода введите exit):")
		fmt.Scan(&text)

		if text == "exit" {
			break
		}
		
		b.WriteString(fmt.Sprintf("%v %v %v\n", i, timeNow, text)) 
		i = i + 1
	}

	fileName := "log.txt"
	
	if err := ioutil.WriteFile(fileName, b.Bytes(), 0400); err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", resultBytes)
}