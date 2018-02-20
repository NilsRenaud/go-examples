package types

import (
	"fmt"
	"time"
)

//Error are just an interface to implement containing only one method : Error()
// Usualy functions returns error which are null if nothing bad happend.
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",	e.When, e.What)
}

func Run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
