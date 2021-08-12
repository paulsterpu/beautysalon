package main

import (
	"fmt"
	"github.com/beautysalon/salon"
	"golang.org/x/sync/semaphore"
)

type Order = salon.Order
type Action = salon.Action

const Man = salon.Man
const Woman = salon.Woman
const Child = salon.Child
const Haircut  = salon.Haircut
const Styling  = salon.Styling
const Shaving  = salon.Shaving
const Manicure  = salon.Manicure

func main()  {
	orders := []Order {
		{
			Actions: []Action{Haircut, Shaving},
			RemainingActions: []Action{Haircut, Shaving},
			ClientType: Man,
			OrderSemaphore: semaphore.NewWeighted(1),
			ID: "0",
		},
		{
			Actions: []Action{Haircut, Shaving},
			RemainingActions: []Action{Haircut, Shaving},
			ClientType: Man,
			OrderSemaphore: semaphore.NewWeighted(1),
			ID: "1",
		},
		{
			Actions: []Action{Haircut, Styling, Manicure},
			RemainingActions: []Action{Haircut, Styling, Manicure},
			ClientType: Woman,
			OrderSemaphore: semaphore.NewWeighted(1),
			ID: "2",
		},
		{
			Actions: []Action{Haircut},
			RemainingActions: []Action{Haircut},
			ClientType: Child,
			OrderSemaphore: semaphore.NewWeighted(1),
			ID: "3",
		},
		{
			Actions: []Action{ Styling, Manicure},
			RemainingActions: []Action{ Styling, Manicure},
			ClientType: Woman,
			OrderSemaphore: semaphore.NewWeighted(1),
			ID: "4",
		},
	}

	fmt.Println("Starting the day...")

	salon.ScheduleOrders(orders)
}