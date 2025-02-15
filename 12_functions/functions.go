package main

import "fmt"

//function -> block of code that performs a specific task
//func name(parameters) return type {}

func add(x int, y int) int {
	return x + y
}

//multiple return
func getLanguages() (string, string, string) {
	return "golang", "python", "java"
}

//function as a parameter
func processIt(fn func(a int) int) int {
	res :=fn(3)
  return res
}

//function as a return type
func getAdder() func(a int) int {
  return func(a int) int {
    return a + 1
  }
}

func main() {
	result := add(10, 20)
	fmt.Println(result)

	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang2, lang3)

	//function in a variable
	fun := func(a int) int {
		return a*a
	}
	sqr := processIt(fun)
	fmt.Println(sqr)

  //function as a return type
  adder := getAdder()
  fmt.Println(adder(10))
}
