package main

import (
	"fmt"
	"os/user"
)

func main() {
	us, _ := user.Current()
	fmt.Println(us.HomeDir)
	fmt.Printf("%T %v", us, us)
}
