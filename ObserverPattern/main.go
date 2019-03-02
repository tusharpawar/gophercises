package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tusharpawar/gophercises/ObserverPattern/consumer"
	"github.com/tusharpawar/gophercises/ObserverPattern/state"
)

const (
	Create = "CREATE"
	Pause  = "PAUSE"
	Resume = "RESUME"
	Stop   = "STOP"
)

func main() {

	stateObj := state.StateMachine{
		State: Create,
	}

	a := consumer.ConA{
		Name: "ConsumerA",
	}

	b := consumer.ConB{
		Name: "ConsumerB",
	}

	//Registering both consumers
	stateObj.Register(&a)

	fmt.Println("")

	stateObj.Register(&b)

	fmt.Println("\nChange the state of StateMachine, Enter \n0 - Exit \n1 - Create \n2 - Pause \n3 - Resume \n4 - Stop")

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, Try again")
			break
		}

		input = strings.Replace(input, "\n", "", -1)

		switch input {
		case "0":
			os.Exit(0)
		case "1":
			stateObj.NotifyConsumers(Create)
		case "2":
			stateObj.NotifyConsumers(Pause)
		case "3":
			stateObj.NotifyConsumers(Resume)
		case "4":
			stateObj.NotifyConsumers(Stop)
		default:
			fmt.Println("Invalid input, try again")
		}
	}

}
