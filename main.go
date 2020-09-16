package main

import (
	"fmt"
	"math"
)

//	- hen price: 150
//	- Egg production formula: (amount * 24) - (30%)
//		hen price / result formula # time left in hours

// DAYHOURS - day hours
const (
	DAYHOURS  = 24
	COMMISION = 0.3
	HENPRICE  = 150
)

func main() {
	var total float64
	amount := 450
	total = float64(amount) * DAYHOURS / 100
	difference := total * COMMISION
	result := total - difference
	totalDays := HENPRICE / result
	totalHours := totalDays * DAYHOURS

	fmt.Println(difference, total)
	fmt.Println("result is :", result)
	fmt.Println("Time to buy a new hen (days):", totalDays)
	fmt.Println("Time to buy a new hen (hours):", totalHours)
	calculateDecimalToTime(totalHours)
}

func calculateDecimalToTime(totalHours float64) {
	var days, hours, minutes, seconds int
	days = int(math.Floor(totalHours) / 24)
	intpart, _ := math.Modf(totalHours)
	hours = int(math.Floor(intpart)) % DAYHOURS
	minutes = int(math.Floor(totalHours*60)) % 60
	seconds = int(math.RoundToEven(totalHours*3600)) % 60

	fmt.Printf("%d days, %02d hours, %02d minutes e %02d seconds", days, hours, minutes, seconds)

}
