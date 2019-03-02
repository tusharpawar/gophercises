package state

import (
	"fmt"

	"github.com/tusharpawar/gophercises/ObserverPattern/consumer"
)

//StateMachine is struct which implements stateMachine interface
type StateMachine struct {
	State        string
	ConsumerList []consumer.Consumer
}

//Register registers any consumer with State
func (s *StateMachine) Register(c consumer.Consumer) {
	if s.ConsumerList == nil {
		s.ConsumerList = make([]consumer.Consumer, 0)
	}

	s.ConsumerList = append(s.ConsumerList, c)
	fmt.Println("consumer is registered with State machine")

	//notify current state to the consumer at the time of registration
	c.Notify(s.State)
}

//NotifyConsumers to notify each consumer who is registered with state
func (s *StateMachine) NotifyConsumers(state string) {
	for _, value := range s.ConsumerList {
		value.Notify(state)
	}
}

//Unregister unregisters any consumer from state
func (s *StateMachine) Unregister(c consumer.Consumer) {
	for i, value := range s.ConsumerList {
		if value == c {
			s.ConsumerList = append(s.ConsumerList[:i], s.ConsumerList[i+1:]...)
		}
	}
	fmt.Println("consumer is unregistered")
}
