package events

type Event interface {
	Type() string
	Execute() error
}
