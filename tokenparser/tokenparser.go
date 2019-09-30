package tokenparser

import (
	"../operations"
	"fmt"
	"strings"
)

/*
* Classily each input line as a token
* A token will have: keyword and argument(s)
* A keyword can be -set alias, -set parameterized alias, etc
* Validation layer:
*       - validate keyword
*       - validate argument(s) for respective keywords
* On successful validation,
*       - Assign the operation to be performed based on the keyword
 */

// parse the token and return the opration that needs to be performed
// along with the operands
func Parse(token string) (operations.Operation, []string) {
	token = removeExtraSpaces(token)
	operation, errorMessage := validateToken(token)
	if operation == operations.RETRY {
		fmt.Println(errorMessage)
	}
	// getting the operands
	tokenSlice := strings.Split(token, " ")
	operands := append(tokenSlice[:0], tokenSlice[1:]...) // removing first element
	return operation, operands
}

// validation tokens and returns the operation to be performed
// Performs: valdate keyword, get operation to perform, validate arguments
// passed to perform operations
func validateToken(token string) (operations.Operation, string) {
	errorMessage := ""
	success := true

	tokenSlice := strings.Split(token, " ")
	keyword := tokenSlice[0]
	success = isValidKeyword(keyword)
	if !success {
		// return retry operation
		errorMessage = "Wrong keyword found! Retry again with correct keyword."
		return operations.RETRY, errorMessage
	}
	op := getOperationToPerform(keyword)
	argsSlice := append(tokenSlice[:0], tokenSlice[1:]...)
	success = isValidArgument(argsSlice, op)
	if !success {
		// return retry operation
		errorMessage = "Arguments to the opreation " + string(op) + " is wrong"
		return operations.RETRY, errorMessage
	}
	return op, errorMessage
}

// with very simple validations, TODO, we can also validate based on colors,
// regex for registraion num, etc
func isValidArgument(args []string, op operations.Operation) bool {

	success := true

	//TODO: write all cases and return from the cases itself

	// ideally this code path should never be reached
	// we should throw exceptions is such cases
	return success
}

func isValidKeyword(keyword string) bool {
	switch keyword {
	case string(operations.ALIAS):
		return true
	case string(operations.ALIAS_P):
		return true
	case string(operations.EXIT):
		return true
		// TODO: implement for other cases aswell
	default:
		return false
	}
}

func getOperationToPerform(keyword string) operations.Operation {
	switch keyword {
	case string(operations.ALIAS):
		return operations.ALIAS
	case string(operations.ALIAS_P):
		return operations.ALIAS_P
	case string(operations.STATUS):
		return operations.STATUS
	case string(operations.EXIT):
		return operations.EXIT
	default:
		return operations.RETRY
	}
}

func removeExtraSpaces(token string) string {
	token = strings.TrimSpace(token)
	token = strings.TrimRight(token, "\n")
	return token
}

func IsExitCommand(input string) bool {
	return strings.EqualFold(strings.TrimRight(input, "\n"), "exit")
}
