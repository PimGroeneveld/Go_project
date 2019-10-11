package main // package to be able to exec file

import (
	"fmt"
)

//sort of the "object"
type person struct {
	name string
	age  int
}

func main() {
	var x int = 5 // "long" way
	y := 10       // short hand

	fmt.Println(x, y)
	fmt.Printf("Hello, World!\n")

	// basic if stuff
	if x > y {
		fmt.Printf("x is bigger than y")
	} else if x < y {
		fmt.Printf("nope x is smaller than y")
	} else {
		fmt.Printf("either equal or something has gone wrong")
	}

	//Basic array stuff
	var array [5]int //index start at 0 like usual
	array[3] = 2345678
	shorthandArray := [5]int{5, 4, 3, 2, 1}
	fmt.Println(array, shorthandArray)

	//solution fixed lenght array: Slices -> no fixed lenght
	slice := []int{435, 5467, 78, 2, 1}
	fmt.Println(slice)
	slice = append(slice, 15) //does not append slice, but returns new one
	fmt.Println(slice)

	// Maps
	// map[type keys]type values in map
	vertices := make(map[string]int)
	vertices["pizza"] = 10
	vertices["fruit"] = 15
	vertices["stroopwafel"] = 7
	fmt.Println(vertices)
	delete(vertices, "pizza")
	fmt.Println(vertices)

	// For loop = only loop in go
	for i := 10; i > 3; i-- {
		fmt.Println(i)
	}
	// for doubles as a while loop
	while := 0
	for while < 5 {
		fmt.Println(while)
		while++
	}

	// need to continue with range
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	// calling func sum
	result := sum(4, 6)
	fmt.Println(result)
}

//func name (params) return type { }
func sum(x int, y int) int {
	return x + y
}

// look into pointers
// errors can be imported and then used as return type (as replacement for Exceptions)
