package main

import (
	"flag"
	"fmt"
)

var (
	typeHen      int
	totalBirds   int
	producedEggs int
)

const (
	RAVENCHICK = iota + 1
	COTTONCHICK
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

var hens map[int]*Hen

func init() {
	hens = make(map[int]*Hen)
	hens[RAVENCHICK] = &Hen{"Raven Chick", 150, 10, 0, 0, 21429}
	hens[COTTONCHICK] = &Hen{"Cotton Chick", 1500, 105, 0, 0, 0}

	flag.IntVar(&typeHen, "t", 1, "Hen type. Default is 1")
	flag.IntVar(&totalBirds, "b", 0, "Total birds from type")
	flag.IntVar(&producedEggs, "p", 0, "Amount of eggs produced by a hen type")

}

func main() {
	flag.Parse()

	if typeHen < RAVENCHICK || totalBirds < 0 || producedEggs < 0 {
		flag.Usage()
		return
	}

	hen := getHenByType(typeHen)
	hen.totalBirds = totalBirds
	hen.laidEggs = producedEggs

	run(hen)
}

func run(hen *Hen) {

	totalHours := calculateTimeProduction(hen)
	fmt.Print("Production Time: ")
	fmt.Println(printTimeCalculated(calculateDecimalToTime(totalHours)))
	remainingHours := calculateTimeRemaining(hen)
	fmt.Print("Remaining Time: ")

	fmt.Println(printTimeCalculated(calculateDecimalToTime(remainingHours)))
}

func getHenByType(typeHen int) *Hen {
	return hens[typeHen]
}
