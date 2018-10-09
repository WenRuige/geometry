package main

import (
	"fmt"
	"math"
)

/*
	平面坐标系两点之间的距离计算公式
	|AB| = sqrt((x1-x2)^2  + (y1-y2)^2)
	地理坐标点两点之间的距离
    |AB| = sqrt((x1-x2)^2  + (y1-y2)^2) * 1e-6
	球面上距离计算公式
    haversin(d/R) = haversin(la2-la1) +cos(la1)cos(la2)haversin(lo2-lo1)
	haversin(theta)   = sin^2 *(theta/2)

*/
// 平面两点之间距离计算公式
func pointtopoint(p1 []float64, p2 []float64) float64 {
	return math.Sqrt(math.Pow(p1[0]-p2[0], 2) + math.Pow(p1[1]-p2[1], 2))
}

// 地图坐标系两点之间距离计算公式 返回距离单位(km)  notice:1e-6 = 10 ^(-6)
func pointtopointgis(p1 []float64, p2 []float64) float64 {
	return math.Sqrt(math.Pow(p1[0]-p2[0], 2)+math.Pow(p1[1]-p2[1], 2)) * 1e-6

}

// 球面距离计算得出的两个点之间的距离
func Distance2(lat1, lon1, lat2, lon2 float64) float64 {
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS
	// calculate
	h := hsin2(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin2(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func hsin2(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func main() {
	p1 := []float64{0, 0}
	p2 := []float64{1, 1}
	res := pointtopoint(p1, p2)

	// 34.816988,113.536474  郑州大学
	// 34.830657,113.550035  河南工业大学   距离为1.97 公里
	p3 := []float64{34.816988, 113.536474}
	p4 := []float64{34.830657, 113.550035}

	res1 := pointtopointgis(p3, p4)
	fmt.Println(res1)
	// 1.925466909608908e-08
	// 1962.4081134222479

	res2 := Distance2(34.816988, 113.536474, 34.830657, 113.550035)
	fmt.Println(res2)
	fmt.Println(res)
}
