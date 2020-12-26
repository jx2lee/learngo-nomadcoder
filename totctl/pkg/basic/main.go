package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// Naked Return funciton
func lenAndUpperNakedReturn(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// defer
func lenAndUpperDefer(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done!")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// for loop
func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// if-else & variable expression
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge > 18 {
		return true
	}
	return false
}

// Switch: variabel type check
func canIDrinkSwitch(age int) bool {
	switch {
	case age+2 < 20:
		return false
	case age+2 >= 20:
		return true
	}
	return false
}

// Struct
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	//fmt.Println(multiply(4, 4))
	//fmt.Println(lenAndUpper("jlee"))
	//repeatMe("lee", "han", "park")

	//totalLength, up := lenAndUpperDefer("jlee")
	//fmt.Println(totalLength, up)

	//result := superAdd(1, 2, 3, 4, 5, 6)
	//fmt.Println(result)

	//fmt.Println(canIDrink(120))
	//fmt.Println(canIDrinkSwitch(10))

	/*	Array
		format: [{size, none}]{type}{value}
		names := [5]string{"lee", "jae", "jun"}
		names := []string{"lee", "jae", "jun"}
		names = append(names, "shin")
		fmt.Println(names)
	*/

	/*	Map
		format: map[key-type]{value-type}{key:value}
		jlee := map[string]string{"name": "jlee", "age": "12"}
		for key := range jlee {
			fmt.Println(key)
		}
	*/

	/*	Struct
		format: Assign struct out of function
		favFood := []string{"kimchi", "ramen"}
		jlee := person{name: "jlee", age: 30, favFood: favFood}
		fmt.Println(jlee)
	*/

}
