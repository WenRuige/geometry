package main

import (
	"fmt"
	"math"
)

// 求最小外接矩形
type rectangle struct {
	maxLat float64
	minLat float64
	maxLng float64
	minLng float64
}

func New() rectangle {
	return rectangle{}
}

// 生成四个值
func (r *rectangle) getMinRectangle(data [][]float64) rectangle {
	maxLat, maxLng := float64(-1<<32), float64(-1<<32)
	minLat, minLng := float64(1<<32), float64(1<<32)
	for _, v := range data {
		maxLat = math.Max(maxLat, v[1])
		maxLng = math.Max(maxLng, v[0])
		minLat = math.Min(minLat, v[1])
		minLng = math.Min(minLng, v[0])

	}
	r.maxLng = maxLng
	r.minLng = minLng
	r.maxLat = maxLat
	r.minLat = minLat
	//r = &rectangle{maxLat: maxLat, minLat: minLat, maxLng: maxLng, minLng: minLat}
	return *r
}

// 生成经纬度
func (r *rectangle) generateLatLng() [][]float64 {
	result := [][]float64{}
	result = append(result, []float64{r.maxLng, r.minLat})
	result = append(result, []float64{r.minLng, r.minLat})
	result = append(result, []float64{r.minLng, r.maxLat})
	result = append(result, []float64{r.maxLng, r.maxLat})
	return result
}

func main() {
	re := [][]float64{{116.365116, 40.099471}, {116.350525, 40.023664}, {116.209934, 40.001182}, {116.186073, 40.104067}}
	obj := New()
	obj.getMinRectangle(re)
	fmt.Println(obj.generateLatLng())

}
