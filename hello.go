package main // package to be able to exec file

import "fmt"

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
	fmt.Println(array)

}
