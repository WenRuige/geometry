package main

import (
	"fmt"

	"os"
	"strings"

	"github.com/mmcloughlin/geohash"
)

// 基于geohash生成一个多边形的四个顶点
func generate(new geohash.Box) [][]float64 {
	result := [][]float64{{}}
	result = append(result, []float64{new.MaxLng, new.MinLat})
	result = append(result, []float64{new.MinLng, new.MinLat})
	result = append(result, []float64{new.MinLng, new.MaxLat})
	result = append(result, []float64{new.MaxLng, new.MaxLat})
	return result
}

// 有一个 mbr矩形
// 初始值的mbr [116.365116 40.001182] [116.186073 40.001182] [116.186073 40.104067] [116.365116 40.104067]
func main() {

	minLat := 40.001182
	minLng := 116.186073
	maxLat := 40.104067
	maxLng := 116.365116
	fmt.Println(maxLat, maxLng, minLat, minLng)

	// 最左侧的点
	originNorthPointGeoHash := geohash.Encode(40.001182, 116.186073)
	// 生成一个BoundBox
	boundingBox := geohash.BoundingBox(originNorthPointGeoHash[:6])
	// 生成的这个是个矩形
	generate(boundingBox)
	// 生成它的邻居
	// 它的邻居为North(北面),NorthEast(东北面),East(东面),SouthEast(东南面),South(南面),SouthWest(西南面),West(西面),NorthWest(西北面)
	neighbors := geohash.Neighbors(originNorthPointGeoHash[:6])
	// 此时会返回8个值
	// 基于北面生成一个BoundingBox
	northBoundingBox := geohash.BoundingBox(neighbors[0])
	// 生成northBoundingBox
	generate(northBoundingBox)
	//循环向右侧生成南面的值,这个值需要小于MAXLAT
	geohash.DecodeCenter(neighbors[0])

	// 初始的Lat

	// x 轴维度的变更

	// y 轴维度的变更
	_, lng := geohash.DecodeCenter(neighbors[0])
	geohash_tmp := neighbors[0]
	for lng < maxLng {
		res := geohash.Neighbors(geohash_tmp)
		//生成四个顶点
		generate(geohash.BoundingBox(res[2]))
		generateJsFile(generate(geohash.BoundingBox(res[2])))
		// 生成一个中心点
		_, lng = geohash.DecodeCenter(res[2])
		geohash_tmp = res[2]
	}
	//最左侧的点,接下来需要向东连续生成
}

func generateJsFile(result [][]float64) {
	str := "["
	for _, v := range result {
		if v == nil {
			continue
		}
		str += fmt.Sprintf("[%v,%v],", v[0], v[1])
	}
	str = "]"

	fmt.Println(str)

}

// 生成文件
func tracefile(str_content string, fileName string) {
	fd, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd_content := strings.Join([]string{str_content, "\n"}, "")
	buf := []byte(fd_content)
	fd.Write(buf)
	fd.Close()
}
