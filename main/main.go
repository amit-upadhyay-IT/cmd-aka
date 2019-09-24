package main

import (
	"../tokenparser"
	"../utils"
	"bytes"
	"fmt"
	"os"
	"os/user"
)

// create a hidden file at /home/username named .aku

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
	fmt.Println(tokens.String())
	operation, operands := tokenparser.Parse(tokens.String())
	fmt.Println(operation, operands)
}

func setup() error {

	// TODO: get file name from some config
	usr, err := user.Current()
	if err != nil {
		fmt.Println("error while getting users homedir")
		return err
	}
	var fullpath = usr.HomeDir + "/.aku"

	if utils.IsFilePresent(fullpath) {
		// file is already present
	} else {
		// create a file with at the specified path
		err := utils.CreateFile(fullpath)
		if err != nil {
			fmt.Println("cannot create the file " + fullpath)
			return err
		}
	}
	return nil
}
