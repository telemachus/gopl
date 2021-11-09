package main

import "fmt"

type Day byte

const (
	Monday Day = 1 << iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
	Tomorrow
)

func main() {
	var requested Day = Monday
	fmt.Println(requested)

	requested = Tomorrow
	fmt.Println(requested)
}

func (d Day) String() string {
	return fmt.Sprintf("%08b", d)
}
