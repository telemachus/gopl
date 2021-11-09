package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

func main() {
	var p = f()
	fmt.Printf("p = %v; *p = %v\n", p, *p)
	p = f()
	fmt.Printf("p = %v; *p = %v\n", p, *p)
	p = f()
	fmt.Printf("p = %v; *p = %v\n", p, *p)
	p = f()
	fmt.Printf("p = %v; *p = %v\n", p, *p)
}
