package golang101

import "fmt"

var zeroValue int
var nothing = "nothing"

const (
	aStr = "a"
	bStr = "b"
)

func printNothing() {
	fmt.Println(nothing)
}

func doSomething(num int) {
	fmt.Println(num)
}

func addSomething(num1 int, num2 int) int {
	return num1 + num2
}

func swap(num1 int, num2 int) (int, int) {
	return num2, num1
}

func swapOneAndTwo() {
	two, one := swap(1, 2)
	fmt.Println(two, one) // prints 2 1

	_, _ = swap(3, 4) // we don't need the returned values
}

func namedReturns() (a, b string) {
	a = aStr
	b = bStr
	return
}
