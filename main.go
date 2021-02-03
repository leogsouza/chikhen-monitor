package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	typeHen      int
	totalBirds   int
	producedEggs int
	goldBalance  float64
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
	name        string
	price       int
	eggsHour    int
	totalBirds  int
	laidEggs    int
	totalToBuy  int
	goldBalance float64
}

var hens map[int]*Hen

func init() {
	hens = make(map[int]*Hen)
	hens[RAVENCHICK] = &Hen{"Raven Chick", 150, 10, 0, 0, 21429, 0}
	hens[COTTONCHICK] = &Hen{"Cotton Chick", 1500, 105, 0, 0, 0, 0}

	flag.IntVar(&typeHen, "th", 1, "Hen type. Default is 1")
	flag.IntVar(&totalBirds, "tb", 0, "Total birds from type")
	flag.IntVar(&producedEggs, "ep", 0, "Amount of eggs produced by a hen type")
	flag.Float64Var(&goldBalance, "gb", 0, "Actual gold balance")

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
	hen.goldBalance = goldBalance
	run(hen)
}

func run(hen *Hen) {

	totalHours := calculateTimeProduction(hen)
	fmt.Print("Production Time: ")
	fmt.Println(printTimeCalculated(calculateDecimalToTime(totalHours)))
	remainingHours := calculateTimeRemaining(hen)
	fmt.Print("Remaining Time: ")
	days, hours, minutes, seconds := calculateDecimalToTime(remainingHours)
	timeEnd := time.Now().Add(time.Duration(seconds) * time.Second).Add(time.Duration(minutes) * time.Minute).Add(time.Duration(hours) * time.Hour).Add(time.Duration(days*HOURSDAY) * time.Hour)
	fmt.Println(printTimeCalculated(days, hours, minutes, seconds))

	fmt.Println("Buy new hen in", timeEnd.Format("02/01/2006 15:04:05"))

}

func getHenByType(typeHen int) *Hen {
	return hens[typeHen]
}
