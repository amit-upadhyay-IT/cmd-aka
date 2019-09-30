package main

import (
	"../operations"
	"fmt"
)

func main() {
	m := operations.GetOperations()
	fmt.Println(m)
}
