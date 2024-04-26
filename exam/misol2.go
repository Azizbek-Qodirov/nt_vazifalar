package main

// import "fmt"

// func main(){
// 	res := convert(5)
// 	fmt.Println(res)
// }

// func convert(n int) string{
// 	ones := [10]string{" ", "bir", "ikki", "uch", "to'rt", "besh", "olti", "yetti", "sakkiz", "to'qqiz"}
// 	tens := [10]string{" ", "o'n", "yigirma", "o'ttiz", "qirq", "ellik", "oltmish", "yetmish", "sakson", "to'qson"}
// 	var result string

// 	birlik := n / 1000
// 	if n >= 1000 {
// 		result += ones[birlik] + " ming "
// 		n %= 1000
// 	}
// 	yuzlik := n / 100
// 	if n >= 100{
// 		result += ones[yuzlik] + " yuz "
// 		n %= 100
// 	}
// 	onlik := n / 10
// 	if n >= 10 {
// 		result += tens[onlik] + " "
// 		n %= 10
// 	}
// 	if n < 10{
// 		result += ones[n]
// 	}
// 	return result

// }

// Done