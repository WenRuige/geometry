package main

import (
	"fmt"
	"math"

	"github.com/geometry/base"
)

/*
	如果想判断一个点是否在线段上需要满足以下两个条件
	1.(Q-P1) * (P2-P1) = 0
	2.Q在以P1,P2为对角线的矩形内
	Notice:向量的叉乘
*/

// 点是否在线段上
func ispointinsegments(Pi base.Point, Pj base.Point, Q base.Point) bool {
	if (Q.X-Pi.X)*(Pj.Y-Pi.Y) == (Pj.X-Pi.X)*(Q.Y-Pi.Y) &&
		math.Min(Pi.X, Pj.X) <= Q.X && Q.X <= math.Max(Pi.X, Pj.X) &&
		math.Min(Pi.Y, Pj.Y) <= Q.Y && Q.Y <= math.Max(Pi.Y, Pj.Y) {
		return true
	}
	return false
}
func main() {
	//q := []float64{1, 1}
	q := base.Point{
		1, 1,
	}
	p1 := base.Point{
		0, 0,
	}
	p2 := base.Point{
		1, 1,
	}
	//p1 := []float64{0, 0}
	//p2 := []float64{1, 1}
	res := ispointinsegments(p1, p2, q)
	fmt.Println(res)
}
