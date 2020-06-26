package presentation

import "fmt"

type StringType string

func doIf() {
	hello := "hello"

	if hello == "hello" {
		fmt.Println(hello)
	} else {
		fmt.Println("bye")
	}
}

func loopSomething() {
	var arr []int
	arr = []int{1, 2, 3}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func forAsWhileLoop() {
	sum := 1

	for sum < 100 {
		sum += 1
	}

	fmt.Println(sum)
}

func loopSomethingWithRange() {
	arr := []int{1, 2, 3}

	for i, num := range arr {
		fmt.Println(num, " at index ", i)
	}
}

func infiniteLoop() {
	for {
		fmt.Println("Twilight Zone")
	}
}

func switchStatement(something int) {
	switch something {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("none of these")
	}
}

func deferSomething() {
	defer fmt.Print(" bye")
	defer fmt.Print(" wazzup")

	fmt.Print("hello")
	// hello wazzup bye
}
