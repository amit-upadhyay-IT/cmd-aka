package app

import (
	"../io"
	"log"
	"os/user"
	"strings"
)

// reading keys from file.
func GetAkaKeys() []string {
	usr, err1 := user.Current()
	if err1 != nil {
		log.Fatal("error while getting users homedir")
	}
	var filepath = usr.HomeDir + "/.aku"
	contents, err2 := io.ReadFile(filepath, true)
	if err2 != nil {
		log.Fatal(err2)
	}
	var keysList []string
	for _, line := range contents {
		// the keys and values are separated by a comma in the file
		pair := strings.Split(line, ",")
		if len(pair) > 0 {
			keysList = append(keysList, pair[0])
		}
	}
	return keysList
}
