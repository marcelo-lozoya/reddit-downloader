package presentation

import "fmt"

type sampleStruct struct {
	something int
}

func InitSampleStruct() sampleStruct {

	return sampleStruct{
		something: 1,
	}

}

func (s sampleStruct) sampleStructFunc1() {
	fmt.Println("im a sample struct!")
	s.sampleStructFunc2()
}

func (s sampleStruct) sampleStructFunc2() {
	fmt.Println("im a sample struct too!")
}

func callSampleStruct() {
	s := InitSampleStruct()
	s.sampleStructFunc1()
}
