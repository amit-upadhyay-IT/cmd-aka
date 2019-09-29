package main

import (
	"fmt"
	"sync"
)

var akaInstance *aka = nil
var once sync.Once

func GetAka() *aka {
	once.Do(func() {
		akaInstance = &aka{}
	})
	return akaInstance
}

type aka struct {
	filepath string
}

func (*aka) dummyMethod() {
	fmt.Println("this is dummy call")
}
