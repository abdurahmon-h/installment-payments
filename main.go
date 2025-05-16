package main

import (
	"fmt"
	"installment-payments/internal/payments"
)

func main() {

	product := "смартфон"
	price := 1000.0
	term := 18
	phone := 992938001313

	total, err := payments.TotalAmount(product, price, term)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	message := fmt.Sprintf("Вы купили %s на %d месяцев. Сумма к оплате: %.2f сомони.",
		product, term, total)

	fmt.Println("Отправка смс на номер:", phone)
	fmt.Println("смс:", message)
}
