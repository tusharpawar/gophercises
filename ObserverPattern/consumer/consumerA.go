package consumer

import "fmt"

//ConA is custom consumer which implements consumer interface
type ConA struct {
	State string
	Name  string
	//other parameters if needed
}

//Notify is method to get the statechange notification
func (c *ConA) Notify(state string) {
	c.State = state
	fmt.Printf("ConA has been notified about state change, state: %s \n", state)
}

//ConA specific methods ...
