package main

import (
	"fmt"
)

func main() {
	var input interface{}
	fmt.Print("Ma'lumot kiriting: \n")
	fmt.Scan(&input)

	switch input.(type) {
	case int:
		fmt.Println("Kiritilgan qiymat int tipida.")
	case float64:
		fmt.Println("Kiritilgan qiymat float tipida.")
	case string:
		fmt.Println("Kiritilgan qiymat string tipida.")
	case bool:
		fmt.Println("Kiritilgan qiymat bool tipida.")
	default:
		fmt.Println("Kiritilgan qiymatning tipi topilmadi.")
	}
}
