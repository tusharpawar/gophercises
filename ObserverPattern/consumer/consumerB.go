package consumer

import "fmt"

//ConB is custom consumer which implements consumer interface
type ConB struct {
	State string
	Name  string
	//other parameters if needed
}

//Notify is method to get the statechange notification
func (c *ConB) Notify(state string) {
	c.State = state
	fmt.Printf("ConB has been notified about state change, state: %s \n", state)
}

//ConB specific methods ...
