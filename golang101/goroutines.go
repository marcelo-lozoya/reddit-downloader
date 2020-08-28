package golang101

import "fmt"

func printSomething() {
	fmt.Println("do something")
}

func doSomethingElse() {
	go printSomething()
}
