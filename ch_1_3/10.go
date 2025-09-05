package main

import (
	"errors"
	"fmt"
)

func main() {
	var psw string
	_, _ = fmt.Scan(&psw)

	err := readPsw(psw)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ok")
}

func readPsw(psw string) error {
	if len(psw) < 6 {
		return errors.New("пароль слишком короткий")
	}

	return nil
}
