package tentsuyutils

import (
	"math"
	"math/rand"
)

//NearCoords tests if (x,y) are "near" (x1, y1) based on radius
func NearCoords(x, y, x1, y1, radius float64) bool {
	//radius := 20.0
	if x <= (x1+radius) && x >= (x1-radius) {
		if y <= (y1+radius) && y >= (y1-radius) {
			return true
		}
	}

	return false
}

//WithinDistance is a simple function to check if x is withing a certain distance of x1
func WithinDistance(x, x1, radius float64) bool {
	if x <= (x1+radius) && x >= (x1-radius) {
		return true
	}
	return false
}

//RandomBetween returns a random int between min and max respectively
//This is just a useful function for many reasons
func RandomBetween(min, max int) int {
	return rand.Intn(max-min) + min
}

//RandomBetweenf is the same as RandomBetween but takes and returns float64 values
func RandomBetweenf(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

//Distance returns the scalar distance between two points
func Distance(x, y, x1, y1 float64) float64 {
	return math.Sqrt(math.Pow(x-x1, 2) + math.Pow(y-y1, 2))
}

//DegreeToRadian takes a given degree and returns the radian value
func DegreeToRadian(degree float64) float64 {
	return degree * math.Pi / 180
}

//RadianToDegree takes the radian value and returns the degree value
func RadianToDegree(radian float64) float64 {
	return radian * 180 / math.Pi
}

//AngleBetween returns the float64 angle of the given 2 sets of x,y coords
func AngleBetween(x1, y1, x2, y2 float64) float64 {
	return math.Atan2(y2-y1, x2-x1)
}

//MakeRangeInt of ints between min and max values
func MakeRangeInt(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

//WhatQuadrant returns an int representing the quadrant of a circle the angle is in
func WhatQuadrant(angle float64) int {

	if math.Sin(angle) > 0 && math.Cos(angle) > 0 {
		return 1
	} else if math.Sin(angle) > 0 && math.Cos(angle) < 0 {
		return 2
	} else if math.Sin(angle) < 0 && math.Cos(angle) < 0 {
		return 3
	} else if math.Sin(angle) < 0 && math.Cos(angle) > 0 {
		return 4
	}
	return -1
}
