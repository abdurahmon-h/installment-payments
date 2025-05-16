package main

import (
	"errors"
	"fmt"
	"strings"
)

var conditions = []int{3, 6, 9, 12, 18, 24}

type ProductConditions struct {
	MainTerms     int
	AddComissions int
}

var productData = map[string]ProductConditions{
	"смартфон":  {9, 3},
	"компьтер":  {12, 4},
	"телевизор": {18, 5},
}

func isValidCondition(term int) bool {
	for _, t := range conditions {
		if t == term {
			return true
		}
	}
	return false
}
func TotalAmount(product string, amount float64, month int) (float64, error) {
	product = strings.ToLower(strings.TrimSpace(product))
	data, match := productData[product]
	if !match {

		return 0, errors.New("Неизвестный продукт")
	}

	if !isValidCondition(month) {
		return 0, errors.New("срок рассрочки должен быть одним из этих: 3, 6, 9, 12, 18, 24 месяцев")
	}

	if month <= data.MainTerms {
		return amount, nil
	}

	extraMonth := 0
	for _, t := range conditions {
		if t > data.MainTerms && t <= month {
			extraMonth++
		}

	}

	commission := float64(extraMonth) * float64(data.AddComissions)
	total := amount + (amount * commission / 100)
	return total, nil
}

func main() {

	product := "смартфон"
	price := 1000.0
	term := 18
	phone := 992938001313

	total, err := TotalAmount(product, price, term)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	message := fmt.Sprintf("Вы купили %s на %d месяцев. Сумма к оплате: %.2f сомони.",
		product, term, total)

	fmt.Println("Отправка смс на номер:", phone)
	fmt.Println("смс:", message)
}
