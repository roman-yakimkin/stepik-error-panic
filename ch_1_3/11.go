package main

import "fmt"

func main() {
	var orderID int
	_, _ = fmt.Scan(&orderID)

	result, err := readOrder(orderID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func readOrder(orderID int) (string, error) {
	if orderID <= 0 {
		return "", fmt.Errorf("некорректный номер заказа: %d", orderID)
	}

	return fmt.Sprintf("Заказ #%d обработан", orderID), nil
}
