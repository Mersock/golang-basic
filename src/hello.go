package main

import "fmt"

//global var
var gVariable int = 5000

//struct
type Books struct {
	title    string
	author   string
	subtitle string
	price    float64
}

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

	// dosomething("kbank")
	// addition(1, 2)
	// anotherFunc()
	// fmt.Println(addition2(5, 2))

	//variadic func
	// summation(1, 2, 3)

	//rescursive func
	// fmt.Println(factorial(5))

	//array
	// var x [5]int
	// x[0] = 10
	// x[1] = 20
	// x[2] = 30
	// x[3] = 40
	// x[4] = 50
	// fmt.Println(x)

	// y := [3]float32{1, 2, 3}
	// fmt.Println(y)

	// fmt.Println(len(x))

	// var total float32
	// for _, value := range y {
	// 	total += value
	// }
	// fmt.Println(total / float32(len(x)))

	//slice
	// x := make([]int, 20)
	// y := []int{1, 2, 3}
	// z := append(y, 4, 5)
	// xx := []int{1, 2, 3, 5}
	// copy(y, x)

	//map
	// x := map[int]string{
	// 	1: "hello",
	// 	2: "hi",
	// }
	// fmt.Println(x[1])

	//closure
	// x := func(x int, y int) int {
	// 	return x + y
	// }
	// fmt.Println(x(1, 2))

	//pointer
	// x := 10
	// fmt.Printf("value is %d\n", x)
	// fmt.Printf("Address x variable %x\n", &x)
	// var p *int
	// p = &x
	// fmt.Printf("pointer p = %x\n", p)
	// *p = 20
	// fmt.Printf("value is %d\n", x)

	//struct
	// var Book1 Books
	// Book1.title = "Hello Go"
	// Book1.author = "Knz"
	// Book1.subtitle = "phumthawan"
	// Book1.price = 222.2
	// fmt.Println(Book1)

	// Book2 := Books{title: "Hi", author: "Hello", price: 123.33}
	// fmt.Println(Book2.price * 2)
	go f(0)
	var input string
	fmt.Scanln(&input)
}

//routine
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

//func basic

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

func summation(arr ...int) {

	var total int
	//no base varable
	for _, n := range arr {
		total += n
	}
	fmt.Println(total)
}

func addition2(a int, b int) int {
	output := a + b
	return output
}

//Variadic function
func summation1(a int, b int) {
	fmt.Println(a + b)
}

func summation2(a int, b int, c int) {
	fmt.Println(a + b + c)
}

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}
