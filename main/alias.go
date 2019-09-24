package main

import (
	"sync"
)

var aliasInstance *alias = nil
var once sync.Once

func GetAlias() *alias {
	once.Do(func() {
		aliasInstance = &alias{}
	})
	return aliasInstance
}

type alias struct {
	filepath string
}

