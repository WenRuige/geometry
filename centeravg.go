package main

import "fmt"

// 计算中心点
func calCenterAvg(scope [][]float64) []float64 {
	sumLng := float64(0)
	sumLat := float64(0)
	for _, v := range scope {
		sumLng += v[0]
		sumLat += v[1]
	}
	return []float64{(sumLng) / float64(len(scope)), (sumLat) / float64(len(scope))}

}

func main() {
	originScope := [][]float64{{116.304233, 39.986398}, {116.449112, 39.992589}, {116.493381, 39.915599}, {116.33068, 39.92224}}
	result:= calCenterAvg(originScope)
	fmt.Println(result)

}
