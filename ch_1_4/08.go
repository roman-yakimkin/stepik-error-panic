package main

import (
	"fmt"
	"strings"
)

func main() {
	var filename string

	_, _ = fmt.Scan(&filename)

	if strings.TrimSpace(filename) == "" {
		fmt.Println("пустое имя файла")
		return
	}

	vals := strings.Split(filename, ".")
	if len(vals) == 0 || vals[len(vals)-1] != "txt" {
		fmt.Println("неверное расширение")
		return
	}

	fmt.Printf("файл принят")
}
