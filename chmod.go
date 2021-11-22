package main

import (
	"fmt"
	_"io"
	"os"
)

func main() {
	
    file, err := os.Create("file.txt")
    if err != nil {
		fmt.Println(err)
	}
    
	file.WriteString(fmt.Sprintf("hello world")) 
	
	defer file.Close()
    
	if err := os.Chmod("file.txt", 0777); err != nil {
        fmt.Println(err)
    }
}
