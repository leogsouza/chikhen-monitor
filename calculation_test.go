package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculationToTime(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		in                     float64
		expD, expH, expM, expS int
	}{
		{2142.85715, 89, 6, 51, 26},
		{42.01681, 1, 18, 1, 1},
	}

	for _, tc := range tests {
		days, hours, minutes, seconds := calculateDecimalToTime(tc.in)
		assert.Equal(days, tc.expD)
		assert.Equal(hours, tc.expH)
		assert.Equal(minutes, tc.expM)
		assert.Equal(seconds, tc.expS)
	}
}

func TestCalculateTimeProduction(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		in       *Hen
		expected float64
	}{
		{&Hen{"Raven Chick", 150, 10, 51, 7000, 0}, 2142.85715},
		{&Hen{"Raven Chick", 150, 10, 1, 0, 0}, 2142.85715},
	}

	for _, tc := range tests {
		assert.Equal(calculateTimeProduction(tc.in), tc.expected)
	}
}

func TestCalculateTimeRemaining(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		in       *Hen
		expected float64
	}{
		{&Hen{"Raven Chick", 150, 10, 51, 7000, 0}, 28.29131652661065},
		{&Hen{"Raven Chick", 150, 10, 1, 0, 0}, 2142.857142857143},
	}

	for _, tc := range tests {
		assert.Equal(calculateTimeRemaining(tc.in), tc.expected)
	}
}

func TestPrintTimeCalculated(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		days, hours, minutes, seconds int
		expected                      string
	}{
		{1, 18, 1, 1, "1 days, 18 hours, 01 minutes and 01 seconds\n"},
		{1, 4, 17, 32, "1 days, 04 hours, 17 minutes and 32 seconds\n"},
	}

	for _, tc := range tests {
		assert.Equal(tc.expected, printTimeCalculated(tc.days, tc.hours, tc.minutes, tc.seconds))
	}
}
