package consumer

//Consumer is an interface for any consumer to implement
type Consumer interface {
	Notify(state string)
}
