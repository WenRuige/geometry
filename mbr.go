package main
//
//import (
//	"math"
//	"fmt"
//	"strings"
//)
//
//// 求最小外接矩形
//type Rectangle struct {
//	MaxLat float64
//	MinLat float64
//	MaxLng float64
//	MinLng float64
//}
//
//func New() Rectangle {
//	return Rectangle{}
//}
//
//// 生成四个值
//func (r *Rectangle) GetMinRectangle(data [][]float64) Rectangle {
//	maxLat, maxLng := float64(-1<<32), float64(-1<<32)
//	minLat, minLng := float64(1<<32), float64(1<<32)
//	for _, v := range data {
//		maxLat = math.Max(maxLat, v[1])
//		maxLng = math.Max(maxLng, v[0])
//		minLat = math.Min(minLat, v[1])
//		minLng = math.Min(minLng, v[0])
//
//	}
//	r.MaxLng = maxLng
//	r.MinLng = minLng
//	r.MaxLat = maxLat
//	r.MinLat = minLat
//	//r = &rectangle{maxLat: maxLat, minLat: minLat, maxLng: maxLng, minLng: minLat}
//	return *r
//}
//
//
///*
//
//[116.403322,39.920255],[116.410703,39.897555],[116.402292,39.892353],[116.389846,39.891365]
//*/
//
//// 生成经纬度
//func (r *Rectangle) generateLatLng() [][]float64 {
//	result := [][]float64{}
//	result = append(result, []float64{r.MaxLng, r.MinLat})
//	result = append(result, []float64{r.MinLng, r.MinLat})
//	result = append(result, []float64{r.MinLng, r.MaxLat})
//	result = append(result, []float64{r.MaxLng, r.MaxLat})
//	return result
//}
//
//func generatesJsFile(result [][]float64) string {
//	str := "["
//	for _, v := range result {
//		if v == nil {
//			continue
//		}
//		str += fmt.Sprintf("[%v,%v],", v[0], v[1])
//	}
//	str = strings.TrimSuffix(str, ",")
//	str += "]"
//	return str
//
//}
//
//func main() {
//	re := [][]float64{{116.304233,39.986398}, {116.449112,39.992589}, {116.493381,39.915599}, {116.33068,39.92224}}
//	obj := New()
//	obj.GetMinRectangle(re)
//	///fmt.Println(obj.generateLatLng())
//	va := generatesJsFile(obj.generateLatLng())
//	fmt.Println(va)
//}
