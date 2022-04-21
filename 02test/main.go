package main

import "fmt"

func main() {
	var s string
	s = f1()
	fmt.Println(&s)

}

func f1() string {
	var str = "hello, world!"
	fmt.Println(&str)
	return str
}
