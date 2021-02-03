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

// Hens types
const (
	RAVEN = iota + 1
	COTTON
	SAPPHIRE
	SPAROW
	CLOVER
	GOLDEN
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
	goldBalance float64
}

var hens map[int]*Hen

func init() {
	hens = make(map[int]*Hen)
	hens[RAVEN] = &Hen{"Raven Chick", 150, 10, 0, 0, 0}
	hens[COTTON] = &Hen{"Cotton Chick", 1500, 105, 0, 0, 0}
	hens[SAPPHIRE] = &Hen{"Sapphire Chick", 7500, 540, 0, 0, 0}
	hens[SPAROW] = &Hen{"Sparow Chick", 15000, 1100, 0, 0, 0}
	hens[CLOVER] = &Hen{"Clover Chick", 150000, 12500, 0, 0, 0}
	hens[GOLDEN] = &Hen{"Golden Chick", 450000, 40000, 0, 0, 0}

	flag.IntVar(&typeHen, "th", 1, "Hen type. Default is 1")
	flag.IntVar(&totalBirds, "tb", 0, "Total birds from type")
	flag.IntVar(&producedEggs, "ep", 0, "Amount of eggs produced by a hen type")
	flag.Float64Var(&goldBalance, "gb", 0, "Actual gold balance")

}

func main() {
	flag.Parse()

	if typeHen < RAVEN || totalBirds < 0 || producedEggs < 0 {
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
