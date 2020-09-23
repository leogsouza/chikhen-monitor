package main

import (
	"fmt"
	"math"
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

	hen := &Hen{"Raven Chick", 150, 10, 48, 4284, 21429}

	totalHours := measureTimeProduction(hen)

	fmt.Print("Production Time: ")
	calculateDecimalToTime(totalHours)
	fmt.Print("Remaining Time: ")
	remainingHours := measureTimeRemaining(hen)
	calculateDecimalToTime(remainingHours)
}

func calculateDecimalToTime(totalHours float64) {
	var days, hours, minutes, seconds int
	days = int(math.Floor(totalHours) / 24)
	intpart, _ := math.Modf(totalHours)
	hours = int(math.Floor(intpart)) % HOURSDAY
	minutes = int(math.Floor(totalHours*60)) % 60
	seconds = int(math.RoundToEven(totalHours*3600)) % 60

	fmt.Printf("%d days, %02d hours, %02d minutes e %02d seconds\n", days, hours, minutes, seconds)

}

func measureTimeProduction(hen *Hen) float64 {

	var total float64
	total = float64(hen.eggsHour*hen.totalBirds) * HOURSDAY / 100
	difference := total * COMMISION
	result := total - difference
	totalDays := float64(hen.price) / result
	totalHours := totalDays * HOURSDAY

	return totalHours
}

func measureTimeRemaining(hen *Hen) float64 {

	return (float64(hen.totalToBuy) - float64(hen.laidEggs)) / (float64(hen.totalBirds) * float64(hen.eggsHour))
}
