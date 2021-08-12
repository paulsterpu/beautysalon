package salon

import (
	"fmt"
	"golang.org/x/sync/semaphore"
)

type ClientType int

const (
	Man ClientType = iota
	Woman
	Child
)

type Order struct {
	Actions[] Action
	RemainingActions[] Action
	CurrentAction Action
	ClientType ClientType
	OrderSemaphore *semaphore.Weighted	// used to make sure we send a client to only 1 chair at a time
	ID string
}

func (order *Order) String() string {
	return fmt.Sprintf("%v, actions: %v; remaining actions %v", order.ClientType, order.Actions, order.RemainingActions)
}

func (client ClientType) String() string {
	switch client {
	case Man:
		return "Man"
	case Woman:
		return "Woman"
	case Child:
		return "Child"
	}

	return ""
}