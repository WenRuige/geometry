package engine

import (
	"strings"

	"github.com/mmcloughlin/geohash"
	"os"
)

//North(北面),NorthEast(东北面),East(东面),SouthEast(东南面),South(南面),SouthWest(西南面),West(西面),NorthWest(西北面)
const (
	NORTH     = iota
	NORTHEAST
	EAST
	SOUTHEAST
	SOUTH
	SOUTHWEST
	WEST
	NORTHWEST
)

// geohash 等级
const (
	GEO_HASH_LAVEL_0 = iota
	GEO_HASH_LAVEL_1
	GEO_HASH_LAVEL_2
	GEO_HASH_LAVEL_3
	GEO_HASH_LAVEL_4
	GEO_HASH_LAVEL_5
	GEO_HASH_LAVEL_6
)

// 地块
type GridInfo struct {
	Position int64       // 位置
	Geohash  int64       // geohash的值
	Level    int64       // geohash的等级
	Scope    [][]float64 // 范围
}

// 地理引擎
type GeographicEngine struct {
	OriginScope [][]float64 // 原始的商圈
	MBRScope    [][]float64
	GridList    []GridInfo
}

// 基于geohash生成一个多边形的四个顶点
func generate(new geohash.Box) [][]float64 {
	result := [][]float64{}
	result = append(result, []float64{new.MaxLng, new.MinLat})
	result = append(result, []float64{new.MinLng, new.MinLat})
	result = append(result, []float64{new.MinLng, new.MaxLat})
	result = append(result, []float64{new.MaxLng, new.MaxLat})
	return result
}

var geographicEngine GeographicEngine

func Dispatch(originScope [][]float64) {
	// 0.判断是否是多边形 是否是凸包

	// 1.生成mbr scope
	rectangle := GetMinRectangle(originScope)
	// 2.递归生成
	// 3.递归时候,应该要计算生成的网格是否在多边形内,以及他们的关系
	GenerateGridList(rectangle.MaxLat, rectangle.MinLng, rectangle.MaxLng, rectangle.MaxLat, GEO_HASH_LAVEL_5)

	GenerateJSString(geographicEngine)
	//fmt.Println(result)
}

// 递归生成网格列表
func GenerateGridList(lat, lng float64, maxLng float64, maxLat float64, level int64) {
	if lat < maxLat {
		return
	}
	// 1.这里的操作是将传入的初始值,计算成bounding box
	originGeoHash := geohash.Encode(lat, lng)
	k := geohash.BoundingBox(originGeoHash[:level])

	gridInfo := GridInfo{
		Scope: generate(k),
	}
	geographicEngine.GridList = append(geographicEngine.GridList, gridInfo)
	// 递归执行
	recursion(lat, lng, maxLng, level)

	boundingBox := ProduceBoundingBox(lat, lng, SOUTH, level)
	lat, lng = boundingBox.Center()
	GenerateGridList(lat, lng, maxLng, maxLat, level)
}

func ProduceBoundingBox(lat, lng float64, direction, level int64) geohash.Box {
	originNorthPointGeoHash := geohash.Encode(lat, lng)
	neighbors := geohash.Neighbors(originNorthPointGeoHash[:level])
	return geohash.BoundingBox(neighbors[direction])
}

func recursion(lat, lng float64, maxLng float64, level int64) {
	if lng > maxLng {
		return
	}
	boundingBox := ProduceBoundingBox(lat, lng, EAST, level)
	gridInfo := GridInfo{
		Scope: generate(boundingBox),
	}
	geographicEngine.GridList = append(geographicEngine.GridList, gridInfo)
	lat, lng = boundingBox.Center()
	recursion(lat, lng, maxLng, level)
}

// 生成文件
func tracefile(str_content string, fileName string) {
	fd, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd_content := strings.Join([]string{str_content, "\n"}, "")
	buf := []byte(fd_content)
	fd.Write(buf)
	fd.Close()
}
