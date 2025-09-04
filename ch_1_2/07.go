package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}

	return a / b, nil
}

func main() {
	var a, b int
	_, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		return
	}

	result, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
