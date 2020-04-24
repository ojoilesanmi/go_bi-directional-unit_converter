package main

import (
	"fmt"
)

type converter struct {
	Feet       float64
	Centimeter float64
	Minutes    float64
	Seconds    float64
	Kilograms  float64
	Pounds     float64
	Degree     float64
	Radian     float64
	Fahrenheit float64
	Celsius    float64
}

func feetToCentimeter(Centimeter float64) float64 {
	return Centimeter * 30.48

}

func centimeterToFeet(Feet float64) float64 {
	return Feet / 30.48
}

func minutesToSeconds(Seconds float64) float64 {
	return Seconds * 60
}

func secondsToMinutes(Minutes float64) float64 {
	return Minutes / 60
}

func main() {
	fmt.Println(feetToCentimeter(12.5))
	fmt.Println(centimeterToFeet(300))
	fmt.Println(minutesToSeconds(430))
	fmt.Println(secondsToMinutes(67))
}
