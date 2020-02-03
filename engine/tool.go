package engine

import (
	"strings"
	"fmt"
	"github.com/mmcloughlin/geohash"
)

// 生成js字符串
func GenerateJSString(geographicEngine GeographicEngine) string {
	str := "["
	for _, v := range geographicEngine.GridList {
		str += generatesJs(v.Scope)
		str += ","
	}
	str += strings.TrimSuffix(str, ",")
	str += "]"
	return str

}
func generatesJs(result [][]float64) string {
	str := "["
	for _, v := range result {
		if v == nil {
			continue
		}
		str += fmt.Sprintf("[%v,%v],", v[0], v[1])
	}
	str = strings.TrimSuffix(str, ",")
	str += "]"
	return str

}

// 多边形重心
func Gravity(polygon [][]float64, n int) []float64 {
	area := float64(0)
	center := []float64{}
	for i := 0; i < n-1; i++ {
		area += (polygon[i][0]*polygon[i+1][1] - polygon[i+1][0]*polygon[i][1]) / 2
		center[0] += (polygon[i][0]*polygon[i+1][1] - polygon[i+1][0]*polygon[i][1]) * (polygon[i][0] + polygon[i+1][0]);
		center[1] += (polygon[i][0]*polygon[i+1][1] - polygon[i+1][0]*polygon[i][1]) * (polygon[i][1] + polygon[i+1][1]);
	}
	area += (polygon[n-1][0]*polygon[0][1] - polygon[0][0]*polygon[n-1][1]) / 2;
	center[0] += (polygon[n-1][0]*polygon[0][1] - polygon[0][0]*polygon[n-1][1]) * (polygon[n-1][0] + polygon[0][0]);
	center[1] += (polygon[n-1][0]*polygon[0][1] - polygon[0][0]*polygon[n-1][1]) * (polygon[n-1][1] + polygon[0][1]);
	center[0] /= 6 * area;
	center[1] /= 6 * area;
	return center
}

func PointCmp(pointA []float64, pointB []float64, center []float64) bool {
	if pointA[0] >= 0 && pointB[0] < 0 {
		return true
	}

	if pointA[0] == 0 && pointB[0] == 0 {
		return pointA[1] > pointB[1]
	}

	det := (pointA[0]-center[0])*(pointB[1]-center[1]) - (pointB[0]-center[0])*(pointA[1]-center[1])

	if det < 0 {
		return true
	}
	if det > 0 {
		return false
	}

	d1 := (pointA[0]-center[0])*(pointA[0]-center[0]) + (pointA[1]-center[1])*(pointA[1]-center[1]);
	d2 := (pointB[0]-center[0])*(pointB[0]-center[1]) + (pointB[1]-center[1])*(pointB[1]-center[1]);
	return d1 > d2;
}

func ClockwiseSortPoints(polygon [][]float64) [][]float64 {
	//计算重心
	center := make([]float64, 2)
	x := float64(0)
	y := float64(0)
	for i := 0; i < len(polygon); i++ {
		x += polygon[i][0];
		y += polygon[i][1];
	}
	center[0] = x / float64(len(polygon))
	center[1] = y / float64(len(polygon))

	//冒泡排序
	for i := 0; i < len(polygon)-1; i++ {
		for j := 0; j < len(polygon)-i-1; j++ {
			if (PointCmp(polygon[j], polygon[j+1], center)) {
				tmp := polygon[j];
				polygon[j] = polygon[j+1];
				polygon[j+1] = tmp;
			}
		}
	}
	return polygon
}

func GenerateGeohash(point []float64, level int) {
	// 生成geohash值
	originGeoHash := geohash.Encode(point[1], point[0])
	fmt.Println(originGeoHash)
	// 得到geohash对应等级的scope
	box := geohash.BoundingBox(originGeoHash[:level])
	result := generatesJsFile(generate(box))
	fmt.Printf("geohash center part:=%v \n", result)
	// 得到geohash的8个neighbors
	str := geohash.Neighbors(originGeoHash[:level])
	for _, v := range str {
		// 获取neighbors的scope
		box := geohash.BoundingBox(v)
		result := generatesJsFile(generate(box))
		fmt.Printf("geohash scope part:=%v \n", result)
	}
}
