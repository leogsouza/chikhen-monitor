package main

import (
	"fmt"
)

//	- hen price: 150
//	- Egg production formula: (amount * 24) - (30%)
//		hen price / result formula # time left in hours

// HOURSDAY - day hours
const (
	HOURSDAY  = 24
	COMMISION = 0.3
	HENPRICE  = 150
)

type Hen struct {
	name       string
	price      int
	eggsHour   int
	totalBirds int
	laidEggs   int
	totalToBuy int
}

func main() {
	run()
}

func run() {
	hen := &Hen{"Raven Chick", 150, 10, 51, 7000, 21429}

	totalHours := calculateTimeProduction(hen)
	fmt.Println(totalHours)
	fmt.Print("Production Time: ")
	fmt.Println(printTimeCalculated(calculateDecimalToTime(totalHours)))
	remainingHours := calculateTimeRemaining(hen)
	fmt.Println(remainingHours)
	fmt.Print("Remaining Time: ")

	fmt.Println(printTimeCalculated(calculateDecimalToTime(remainingHours)))
}
