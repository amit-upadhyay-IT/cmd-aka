package operations

import (
	"../utils/enums"
	"../utils/io"
	"log"
	"runtime"
)

/*
* this is enum for representing various opreations that a user can perform in
* our application, eg: set alias, set alias -p, etc
 */

type Operation string

const (
	ALIAS   Operation = "-a"    // normal alias
	ALIAS_P Operation = "-ap"   // parameterized alias
	ALIAS_D Operation = "-d"    // to delete both type of alias(i.e. normal and parameterized)
	ALIAS_U Operation = "-u"    // to update the keys or values for an alias
	STATUS  Operation = "-s"    // to see all the alias added
	RETRY   Operation = "retry" // incase wrong input is entered by user
	HELP    Operation = "-h"    // show possible commands with details

	PERFORM Operation = "p" // when user will enter command to perform task,
	// the parser should be smart enough to figure out if the user is trying
	// to perform a valid operation, else parser should return RETRY flag.
)

// returns a list of enum values
// Algo:
// get the current file dir location and read the content, use enumutils to get values
func GetOperations() []string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to recover the file name")
		// this gives a call to  os.Exit(1) after printing the message
	}
	lines, err := io.ReadFile(file, true)
	if err != nil {
		log.Fatal(err)
	}
	// this regex can also be used to get constant var name
	//regString := "(^[a-zA-Z0-9_]+) Original([^()]*)"
	enumsMap, err := enums.GetEnumsMap(lines)
	if err != nil {
		log.Fatal(err)
	}
	return enumsMap["Operation"]
}
