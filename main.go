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
		if arr[i] == 73 && x < 3 {
			x = x + 1
		}
		if arr[i] == 86 && x == 0 {
			x = 5
		}
		if arr[i] == 86 && x == 1 {
			x = 4
		}
		if arr[i] == 73 && x >= 5 {
			x += 1
		}
		if arr[i] == 88 && x == 0 {
			x = 10
		}
		if arr[i] == 88 && x == 1 {
			x = 10 - 1
		}
	}
	return x
}

//func recursivRoman(x int, rez string) string{
//	if()
//}

func toRoman(x int) string {
	rez := ""
	element := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	// ASCII I(1) - 73, V(5) - 86, X(10) - 88, L(50) - 76, C(100) - 67, D(500) - 68, M(1000) - 77
	for i := 1000; x != 0; i -= 1 {
		if x >= i && "" != element[i] {
			for n := x / i; n != 0; n-- {
				rez += element[i]
				fmt.Println(rez)
				x -= i
			}
		}
	}
	return rez
}

func romanNum(arr []string) {
	x := romanToNorm(arr[0])
	y := romanToNorm(arr[2])
	fmt.Println(x, " ", arr[1], " ", y)
	switch arr[1] {
	case "*":
		fmt.Println(toRoman(x * y))
	case "/":
		fmt.Println(toRoman(x / y))
	case "+":
		fmt.Println(toRoman(x + y))
	case "-":
		fmt.Println(toRoman(x - y))
	default:
		fmt.Println("Y have mistake with operator")
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

	if x > 10 {
		fmt.Println("Y give x more than 10")
		os.Exit(0)
	}
	if y > 10 {
		fmt.Println("Y give y more than 10")
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

func read() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("l am calculator?:")

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	//fmt.Printf("%s\n", text)
	//fmt.Printf("%q\n", strings.Split(text, " "))
	texts := strings.Split(text, " ")

	textX := texts[0]
	textY := texts[2]

	if len(texts) != 3 {
		fmt.Println("Y problem with numbers of inputs")
		os.Exit(0)
	}

	//ASCII I - 73, V - 86, X - 88
	if (textX[0] == 73 || textX[0] == 86 || textX[0] == 88) && (textY[0] == 73 || textY[0] == 86 || textY[0] == 88) {
		romanNum(texts)
	} else {
		normalNum(texts)
	}
}

func main() {

	read()

}
