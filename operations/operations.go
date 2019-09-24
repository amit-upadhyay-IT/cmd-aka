package operations

import ()

/*
* this is enum for representing various opreations that a user can perform in
* our application, eg: set alias, set alias -p, etc
 */

type Operation string

const (
	ALIAS        Operation = "alias"
	ALIAS_P      Operation = "alias_parameterized"
	DELETE_ALIAS Operation = "delete_alias" // to delete both type of alias(i.e. normal and parameterized)
	UPDATE_ALIAS Operation = "update_alias" // to update the keys or values for an alias
	STATUS       Operation = "status"       // to see all the alias added
	EXIT                   = "exit"         // should not be required
	RETRY                  = "retry"        // incase wrong input is entered by user
)
