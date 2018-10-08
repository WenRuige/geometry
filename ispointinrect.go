package main

import "fmt"

/*
	点是否在矩形内
	延伸:如何判断一个多边形是矩形
*/

func ispointinrect(p []float64, p1 []float64, p2 []float64, p3 []float64, p4 []float64, ) bool {
	return getCross(p1, p2, p)*getCross(p3, p4, p) >= 0 && getCross(p2, p3, p)*getCross(p4, p1, p) >= 0
}

func getCross(p1 []float64, p2 []float64, p []float64) float64 {
	return (p2[0]-p1[0])*(p[1]-p1[0]) - (p[0]-p1[0])*(p2[1]-p1[1])
}

func main() {
	p1 := []float64{0, 5}
	p2 := []float64{0, 0}
	p3 := []float64{5, 0}
	p4 := []float64{5, 5}
	p := []float64{0, 0}
	res := ispointinrect(p, p1, p2, p3, p4)
	fmt.Println(res)
}
