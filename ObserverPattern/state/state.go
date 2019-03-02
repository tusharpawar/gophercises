package state

import "github.com/tusharpawar/gophercises/ObserverPattern/consumer"

//State is interface which every State has to implement
type State interface {
	Register(c consumer.Consumer)
	Unregister(c consumer.Consumer)
	NotifyConsumers(state string)
}
