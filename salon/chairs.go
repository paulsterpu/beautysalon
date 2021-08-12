package salon

import (
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

// initialize the chairs
var firstChair = FirstChair{Chair: Chair {ChairSemaphore: semaphore.NewWeighted(1)}, Actions: []Action{Haircut, Styling}}
var secondChair = SecondChair{Chair: Chair {ChairSemaphore: semaphore.NewWeighted(1)}, Actions: []Action{Shaving}}
var thirdChair = ThirdChair{Chair: Chair {ChairSemaphore: semaphore.NewWeighted(1)}, Actions: []Action{Manicure}}

type Chair struct {
	ChairSemaphore *semaphore.Weighted
	// additional general info about chairs
}

type FirstChair struct {
	Chair
	Actions[] Action
	// additional info about FirstChair
}

type SecondChair struct {
	Chair
	Actions[] Action
	// additional info about SecondChair
}

type ThirdChair struct{
	Chair
	Actions[] Action
	// additional info about ThirdChair
}

type FirstChairActions interface {
	Haircut(order Order)
	Styling(order Order)
}

type SecondChairActions interface {
	Shaving(order Order)
}

type ThirdChairActions interface {
	Manicure(order Order)
}

func (chair Chair) ChairMethod() bool {
	return true
}

func (chair *FirstChair) Haircut(order *Order) {
	go func() {
		fmt.Println(fmt.Sprintf("FirstChair starting haircut for order (%v, %v)", order.ClientType, order.ID))
		time.Sleep(10 * time.Second)
		fmt.Println(fmt.Sprintf("FirstChair finished haircut for order (%v, %v)", order.ClientType, order.ID))

		// remove the action from the remaining actions
		actionIndex := getActionIndex(order.RemainingActions, Haircut)
		if actionIndex >= 0 {
			order.RemainingActions = append(order.RemainingActions[:actionIndex], order.RemainingActions[actionIndex+1:]...)
		}

		// release the chair and the order(client)
		order.OrderSemaphore.Release(1)
		chair.ChairSemaphore.Release(1)
	}()
}

func (chair *FirstChair) Styling(order *Order) {
	go func() {
		fmt.Println(fmt.Sprintf("FirstChair starting styling for order (%v, %v)", order.ClientType, order.ID))
		time.Sleep(10 * time.Second)
		fmt.Println(fmt.Sprintf("FirstChair finished styling for order (%v, %v)", order.ClientType, order.ID))

		actionIndex := getActionIndex(order.RemainingActions, Styling)
		if actionIndex >= 0 {
			order.RemainingActions = append(order.RemainingActions[:actionIndex], order.RemainingActions[actionIndex+1:]...)
		}

		order.OrderSemaphore.Release(1)
		chair.ChairSemaphore.Release(1)
	}()
}

func (chair *SecondChair) Shaving(order *Order) {
	go func() {
		fmt.Println(fmt.Sprintf("SecondChair starting shaving for order (%v, %v)", order.ClientType, order.ID))
		time.Sleep(10 * time.Second)
		fmt.Println(fmt.Sprintf("SecondChair finished shaving for order (%v, %v)", order.ClientType, order.ID))

		actionIndex := getActionIndex(order.RemainingActions, Shaving)
		if actionIndex >= 0 {
			order.RemainingActions = append(order.RemainingActions[:actionIndex], order.RemainingActions[actionIndex+1:]...)
		}

		order.OrderSemaphore.Release(1)
		chair.ChairSemaphore.Release(1)
	}()
}

func (chair *ThirdChair) Manicure(order *Order) {
	go func() {
		fmt.Println(fmt.Sprintf("ThirdChair starting manicure for order (%v, %v)", order.ClientType, order.ID))
		time.Sleep(10 * time.Second)
		fmt.Println(fmt.Sprintf("ThirdChair finished manicure for order (%v, %v)", order.ClientType, order.ID))

		actionIndex := getActionIndex(order.RemainingActions, Manicure)
		if actionIndex >= 0 {
			order.RemainingActions = append(order.RemainingActions[:actionIndex], order.RemainingActions[actionIndex+1:]...)
		}

		order.OrderSemaphore.Release(1)
		chair.ChairSemaphore.Release(1)
	}()
}
