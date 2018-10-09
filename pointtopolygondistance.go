package main

import (
	"math"
)

/*
	点到一个多边形的距离
*/

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func pointtopolygondistance(lat1, lng1, lat2, lng2, lat3, lng3 float64) float64 {
	a := Distance(lat1, lng1, lat2, lng2)
	b := Distance(lat2, lng2, lat3, lng3)
	c := Distance(lat1, lng1, lat3, lng3)
	if b*b >= (c*c + a*a) {
		return c
	}
	if c*c >= (b*b + a*a) {
		return b
	}
	l := (a + b + c) / 2
	s := math.Sqrt(l * (l - a) * (l - b) * (l - c))
	return 2 * s / a
}
