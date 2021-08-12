package salon

import (
	"fmt"
	"reflect"
	"sync"
)

var ordersWG sync.WaitGroup

func scheduleActions(order *Order, ongoingOrders chan *Order) {

	orderInProgress := false

	if len(order.RemainingActions) == 0 {
		ordersWG.Done()
		fmt.Println(fmt.Sprintf("Order %v finished...", order.ID))
		return
	}

	// if the order has an action in progress skip it and put it back in the queue
	if order.OrderSemaphore.TryAcquire(1) == false {
		ongoingOrders <- order
		return
	}

	// iterate over the remaining actions and see which chair supports the action and if it's free
	for _, action := range order.RemainingActions {
		if supportsAction(firstChair.Actions, action) && firstChair.ChairSemaphore.TryAcquire(1) {
			reflect.ValueOf(&firstChair).MethodByName(fmt.Sprintf("%v", action)).Call([]reflect.Value{reflect.ValueOf(order)})
			orderInProgress = true
			break
		}
		if supportsAction(secondChair.Actions, action) && secondChair.ChairSemaphore.TryAcquire(1) {
			reflect.ValueOf(&secondChair).MethodByName(fmt.Sprintf("%v", action)).Call([]reflect.Value{reflect.ValueOf(order)})
			orderInProgress = true
			break
		}
		if supportsAction(thirdChair.Actions, action) && thirdChair.ChairSemaphore.TryAcquire(1) {
			reflect.ValueOf(&thirdChair).MethodByName(fmt.Sprintf("%v", action)).Call([]reflect.Value{reflect.ValueOf(order)})
			orderInProgress = true
			break
		}
	}

	if orderInProgress == false {
		// if the order did not start because there is no available chair yet for the remaining actions then release the semaphore for it
		order.OrderSemaphore.Release(1)
	}

	// put the order back in queue to execute remaining actions
	ongoingOrders <- order
}

func ScheduleOrders(orders []Order) {
	ongoingOrders := make(chan *Order, len(orders))

	// initial scheduling
	for i := 0; i < len(orders); i++ {
		fmt.Println("Scheduling ", i)
		ordersWG.Add(1)
		ongoingOrders <- &orders[i]
	}

	go func() {
		for order := range ongoingOrders  {
			scheduleActions(order, ongoingOrders)
		}
	}()

	ordersWG.Wait()

	fmt.Println("All orders processed...")
}
