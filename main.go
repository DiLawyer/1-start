package main

import (
	"fmt"
)

const (
	USD_EUR         = 0.86
	USD_RUB         = 80.0
	EUR_TO_RUB_RATE = USD_RUB / USD_EUR
)

func main() {
	fmt.Println(USD_EUR, USD_RUB, EUR_TO_RUB_RATE)

}
