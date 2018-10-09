package main

import "fmt"

/*
	射线法判断点是否在多边形内
*/

func euqal(x float64, y float64) bool {
	v := x - y
	const delta float64 = 1e-6
	if v < delta && v > -delta {
		return true
	}
	return false

}

func little(x float64, y float64) bool {
	if euqal(x, y) {
		return false
	}
	return x < y
}
func little_equal(x float64, y float64) bool {
	if euqal(x, y) {
		return true
	}
	return x < y
}
func InPolygon(point []float64, vertices [][]float64) bool {
	x := point[0]
	y := point[1]
	sz := len(vertices)
	is_in := false

	for i := 0; i < sz; i++ {
		j := i - 1
		if i == 0 {
			j = sz - 1
		}
		vi := vertices[i]
		vj := vertices[j]

		xmin := vi[0]
		xmax := vj[0]
		if xmin > xmax {
			t := xmin
			xmin = xmax
			xmax = t
		}
		ymin := vi[1]
		ymax := vj[1]
		if ymin > ymax {
			t := ymin
			ymin = ymax
			ymax = t
		}
		// i//j//aixs_x
		if euqal(vj[1], vi[1]) {
			if euqal(y, vi[1]) && little_equal(xmin, x) && little_equal(x, xmax) {
				return true
			}
			continue
		}

		xt := (vj[0]-vi[0])*(y-vi[1])/(vj[1]-vi[1]) + vi[0]
		if euqal(xt, x) && little_equal(ymin, y) && little_equal(y, ymax) {
			// on edge [vj,vi]
			return true
		}
		if little(x, xt) && little_equal(ymin, y) && little(y, ymax) {
			is_in = !is_in
		}

	}
	return is_in
}

func main() {
	verts := [][]float64{{2, 1}, {8, 1}, {8, 6}, {6, 6}, {6, 3}, {4, 3}, {4, 5}, {2, 5}}
	fmt.Println(InPolygon([]float64{5, 4}, verts))
	fmt.Println(InPolygon([]float64{6, 6}, verts))
	fmt.Println(InPolygon([]float64{6, 5}, verts))
	fmt.Println(InPolygon([]float64{7, 4}, verts))
	fmt.Println(InPolygon([]float64{8, 4}, verts))
	fmt.Println(InPolygon([]float64{3, 3}, verts))
	fmt.Println(InPolygon([]float64{4, 3}, verts))
}
