package main

import (
	"fmt"
	"math"
)

func calculateDecimalToTime(totalHours float64) (days, hours, minutes, seconds int) {
	days = int(math.Floor(totalHours) / 24)
	intpart, _ := math.Modf(totalHours)
	hours = int(math.Floor(intpart)) % HOURSDAY
	minutes = int(math.Floor(totalHours*60)) % 60
	seconds = int(math.RoundToEven(totalHours*3600)) % 60
	return
}

func printTimeCalculated(days, hours, minutes, seconds int) string {
	return fmt.Sprintf("%d days, %02d hours, %02d minutes and %02d seconds\n", days, hours, minutes, seconds)
}

func calculateTimeProduction(hen *Hen) float64 {

	var total float64
	total = float64(hen.eggsHour*hen.totalBirds) * HOURSDAY / 100
	difference := total * COMMISION
	result := total - difference
	totalDays := float64(hen.price) / result
	totalHours := totalDays * HOURSDAY

	return math.Ceil(totalHours*100000) / 100000
}

func calculateTimeRemaining(hen *Hen) float64 {
	var production float64
	production = float64(hen.totalBirds) * float64(hen.eggsHour)
	remaining := hen.totalToBuy - hen.laidEggs
	return float64(remaining) / production
}
