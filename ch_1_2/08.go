package main

import (
	"errors"
	"fmt"
)

func main() {
	var id int
	_, _ = fmt.Scan(&id)

	result, err := readData(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func readData(id int) (string, error) {
	if id <= 0 {
		return "", errors.New("неверный идентификатор")
	}

	return fmt.Sprintf("Данные №%d", id), nil
}
