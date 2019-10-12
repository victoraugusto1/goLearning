package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "This is a test."
	var version float32 = 1.1
	var integerNumber int
	var inferred = "The type of this variable was not specified"
	var inferredFloat = 2.5
	simpleVariable := "See? It's simple."
	fmt.Println("Hello, world.", name)
	fmt.Println("This is the", version, "version.")
	fmt.Println("The Go language has some interesting types. The initial value of an int is", integerNumber)
	fmt.Println("Some types are inferred in Golang. Example:", inferred, ", but the type is", reflect.TypeOf(inferred))
	fmt.Println("If you don't specify the type of a float variable, Golang will assume that the type you want is", reflect.TypeOf(inferredFloat))
	fmt.Println("You don't even have to use the 'var' word to declare variables. Just use := and it's done.", simpleVariable)
	fmt.Println("Anyway, these are the options for this simple program:")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Watch logs")
	fmt.Println("3 - Exit program")

	var selectedOption int

	fmt.Scanf("%d", &selectedOption)

	fmt.Println("The", selectedOption, "has been choosed.")
	fmt.Println("The memory address of the variable that allocates the selected option is", &selectedOption)
	fmt.Println(
		"There's a function simpler than Scanf, which does not require the information of the type of the value that is being filled. The function is simply Scan.")
	var newValue int
	fmt.Println("Please, type a new value:")
	fmt.Scan(&newValue)
	fmt.Println("You typed the", newValue, "int value. If you have typed something different from an int, the showed value will be 0.")
}
