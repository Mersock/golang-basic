package main

import "fmt"

//global var
var gVariable int = 5000

func main() {
	/*
		Initial
	*/
	//fmt.Println("xxxx")
	//fmt.Print("hello, world\n")

	/*
		Variable
	*/
	//var age int = 200
	//age2 := 5000000.22332
	//text := "This is new text"
	//num1, num2 := 323, 22
	//fmt.Println("Variable integer is", age2)
	//fmt.Println("Text common now is", text)
	//fmt.Println("Text number one", num1)
	//fmt.Println("Text number two", num2

	/*
		number and string
	*/
	// number1 := 5
	// number2 := 4
	// p1 := "knz"
	// p2 := "phumthawan"
	//operation
	// fmt.Println(number1 + number2)
	// fmt.Println(number1 - number2)
	// fmt.Println(number1 * number2)
	// fmt.Println(number1 / number2)

	//concat string
	// p3 := p1 + p2
	// fmt.Println(p3)
	//ascii format
	// fmt.Println(p3[1])

	// fmt.Println(p3[0:])
	// fmt.Println(p3[1:3])

	// isEmpty := true
	// isJumping := false
	// fmt.Println(isEmpty)
	// fmt.Println(isJumping)
	// compareNum := 5 >= 3
	// fmt.Println(compareNum)

	//print global var
	// fmt.Println(gVariable)
	// IVariable := 40
	// fmt.Println(IVariable)
	// anotherFunc()

	// input
	// fmt.Print("Input Number: ")
	// var input float64
	// fmt.Scanf("%f", &input)

	// condition := input > 2
	// if condition {
	// 	fmt.Println("Worked!!!")
	// } else {
	// 	fmt.Println("not Worked!!!")
	// }

	// if 6 < 3 || 8 < 5 {
	// 	fmt.Println("Worked!!!")
	// } else {
	// 	fmt.Println("Not Worked!!!")
	// }

	// switch input {
	// case 0:
	// 	fmt.Println("Zero")
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")

	// default:
	// 	fmt.Println("Unknown")
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("kkkk")
	// }
	// n := 1
	// for n <= 200 {
	// 	fmt.Println("xxxx")
	// 	n++
	// }

	// for a := 1; a <= 100; a++ {
	// 	if a%2 == 0 {
	// 		fmt.Println(a, "--")
	// 	} else {
	// 		fmt.Println(a, "-")
	// 	}
	// }

	dosomething("kbank")
	addition(1, 2)
	anotherFunc()
	fmt.Println(addition2(5, 2))
}

func anotherFunc() {
	IVariable := 20
	fmt.Println(IVariable)
	fmt.Println(gVariable)
}

func dosomething(str string) {
	fmt.Println(str)
}

func addition(a int, b int) {
	fmt.Println(a + b)
}

func addition2(a int, b int) int {
	output := a + b
	return output
}
