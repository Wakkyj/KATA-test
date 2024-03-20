package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func romanToNorm(arr string) int {
	x := 0
	// ASCII I - 73, V - 86, X - 88
	for i := 0; i < len(arr); i++ {
		if arr[i] == 77 { // == M(1000)
			if x < 1000 && x != 0 {
				x = 1000 - x
			} else {
				x += 1000
			}
		}
		if arr[i] == 68 { // == D(500)
			if x < 500 && x != 0 {
				x = 500 - x
			} else {
				x += 500
			}
		}
		if arr[i] == 67 { // == C(100)
			if x < 100 && x != 0 {
				x = 100 - x
			} else {
				x += 100
			}
		}
		if arr[i] == 76 { // == L(50)
			if x < 50 && x != 0 {
				x = 50 - x
			} else {
				x += 50
			}
		}
		if arr[i] == 88 { // == X(10)
			if x < 10 && x != 0 {
				x = 10 - x
			} else {
				x += 10
			}
		}
		if arr[i] == 86 { // == V(5)
			if x < 5 && x != 0 {
				x = 5 - x
			} else {
				x += 5
			}
		}
		if arr[i] == 73 { // == I(1)
			if false {

			} else {
				x += 1
			}
		}
	}
	return x
}

func toRoman(x int) string {
	rez := ""
	for x != 0 {
		if x >= 90 && x <= 130 {
			if x == 100 {
				x -= 100
				rez += "C"
			} else {
				x -= 90
				rez += "XC"
			}
		}
		if x >= 40 && x < 90 {
			if x < 50 {
				x -= 40
				rez += "XL"
			} else {
				x -= 50
				rez += "L"
			}
		}
		if x >= 9 && x < 40 {
			if x == 9 {
				x -= 9
				rez += "IX"
			} else {
				x -= 10
				rez += "X"
			}

		}
		if x >= 4 && x < 9 {
			if x == 4 {
				x -= 4
				rez += "IV"
			} else {
				x -= 5
				rez += "V"
			}
		}
		if x > 0 && x < 4 {
			x -= 1
			rez += "I"
		}
	}
	return rez
}

func romanNum(arr []string) {
	x := romanToNorm(arr[0])
	//fmt.Println(x)
	y := romanToNorm(arr[2])
	//fmt.Println(y)
	if x > 10 || x < 1 || y > 10 || y < 1 {
		fmt.Println("Принимаются числа от I до X")
		os.Exit(0)
	}
	//fmt.Println(y)
	switch arr[1] {
	case "*":
		fmt.Println(toRoman(x * y))
	case "/":
		fmt.Println(toRoman(x / y))
	case "+":
		fmt.Println(toRoman(x + y))
	case "-":
		if (x - y) == 0 {
			fmt.Println("Ошибка, нет нуля в римских числах")
		}
		if y > x {
			fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
			os.Exit(0)
		}
		fmt.Println(toRoman(x - y))

	default:
		fmt.Println("Ошибка с операторами")
	}
}

func normalNum(arr []string) {
	x, errX := strconv.Atoi(arr[0])
	y, errY := strconv.Atoi(arr[2])

	if errX != nil {
		log.Fatal(errX)
	}
	if errY != nil {
		log.Fatal(errY)
	}

	if x > 10 || x < 1 || y > 10 || y < 1 {
		fmt.Println("Принимаются числа от 1 до 10")
		os.Exit(0)
	}

	switch arr[1] {
	case "*":
		fmt.Println(x * y)
	case "/":
		fmt.Println(x / y)
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	default:
		fmt.Println("Y have mistake with operator")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	//fmt.Printf("%s\n", text)
	//fmt.Printf("%q\n", strings.Split(text, " "))
	texts := strings.Split(text, " ")

	if len(texts) != 3 {
		fmt.Println("Вывод ошибки, так как формат математической операции " +
			"не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		os.Exit(0)
	}

	textX := texts[0]
	textY := texts[2]

	if ((textX[0] == 73 || textX[0] == 86 || textX[0] == 88) && (textY[0] <= 57 && textY[0] >= 48)) ||
		((textY[0] == 73 || textY[0] == 86 || textY[0] == 88) && (textX[0] <= 57 && textX[0] >= 48)) {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		os.Exit(0)
	}

	//ASCII I - 73, V - 86, X - 88, L -  76
	if (textX[0] == 73 || textX[0] == 86 || textX[0] == 88 || textX[0] == 76) && (textY[0] == 73 || textY[0] == 86 || textY[0] == 88 || textY[0] == 76) {
		romanNum(texts)
	} else {
		normalNum(texts)
	}
}
