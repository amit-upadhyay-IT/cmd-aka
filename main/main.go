package main

import (
	"../operations"
	"../tokenparser"
	"../utils/io"
	"bytes"
	"fmt"
	"os"
	"os/user"
)

// create a hidden file at /home/username named .aku
func setup() error {

	// TODO: get file name from some config
	usr, err := user.Current()
	if err != nil {
		fmt.Println("error while getting users homedir")
		return err
	}
	var fullpath = usr.HomeDir + "/.aku"

	if io.IsFilePresent(fullpath) {
		// file is already present
	} else {
		// create a file with at the specified path
		err := io.CreateFile(fullpath)
		if err != nil {
			fmt.Println("cannot create the file " + fullpath)
			return err
		}
	}
	return nil
}

func main() {

	err := setup()
	if err != nil {
		os.Exit(1)
	}

	programArgs := os.Args
	var tokens bytes.Buffer
	if len(programArgs) > 1 {
		// read command line args
		for i := 1; i < len(os.Args)-1; i++ {
			tokens.WriteString(programArgs[i] + ",")
		}
		tokens.WriteString(programArgs[len(programArgs)-1])
	}
	var token string = tokens.String()
	operation, operands := tokenparser.Parse(token)
	performOperations(operation, operands)
}

func performOperations(operation operations.Operation, operands []string) {

	// create instance of alias
	alias := GetAka()
	alias.dummyMethod()

	switch operation {
	case operations.ALIAS:
		fmt.Println("ALIAS is to be performed")
		fmt.Println(operation, operands)
		// some operation
	case operations.ALIAS_U:
		// some operations
	default:
		fmt.Println("not implemented yet")
		// or better throw error/ panic the program here
	}

}
