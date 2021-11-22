package main

import (
	"os"
	"fmt"
	"time"
)

func main() {

	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	defer file.Close()
	
	text := ""
	i := 1
	
	for {
		timeNow := time.Now().Format("2006-01-02 15:04:05") // 🤯 
		
		fmt.Println("Введите строку (для выхода введите exit):")
		fmt.Scan(&text)

		if text == "exit" {
			break
		}
		
		file.WriteString(fmt.Sprintf("%v %v %v\n", i, timeNow, text)) 
		i = i + 1
	}

	
}


