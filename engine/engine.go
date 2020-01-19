package engine

import (
	"strings"

	"github.com/mmcloughlin/geohash"
	"fmt"
	"os"

	"errors"
	"math"
	"strconv"
)

// North(北面),NorthEast(东北面)
// East(东面),SouthEast(东南面)
// South(南面),SouthWest(西南面)
// West(西面),NorthWest(西北面)
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
	GEO_HASH_LEVEL_0 = iota
	GEO_HASH_LEVEL_1
	GEO_HASH_LEVEL_2
	GEO_HASH_LEVEL_3
	GEO_HASH_LEVEL_4
	GEO_HASH_LEVEL_5
	GEO_HASH_LEVEL_6
)

type Point struct {
	X float64
	Y float64
}

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
	MBRScope    Rectangle
	GridList    []GridInfo // 地块信息
}

// 基于geohash生成一个多边形的四个顶点
func generate(new geohash.Box) [][]float64 {
	result := [][]float64{}
	result = append(result, []float64{Decimal(new.MinLng), Decimal(new.MaxLat)})
	result = append(result, []float64{Decimal(new.MaxLng), Decimal(new.MaxLat)})
	result = append(result, []float64{Decimal(new.MaxLng), Decimal(new.MinLat)})
	result = append(result, []float64{Decimal(new.MinLng), Decimal(new.MinLat)})

	//result = append(result, []float64{new.MinLng, new.MaxLat})
	//result = append(result, []float64{new.MinLng, new.MinLat})
	//result = append(result, []float64{new.MaxLng, new.MinLat})
	//result = append(result, []float64{new.MaxLng, new.MaxLat})
	return result
}

var geographicEngine GeographicEngine

func Dispatch(originScope [][]float64) {
	// 0.判断是否是多边形 是否是凸包

	geographicEngine.OriginScope = originScope
	// 1.生成mbr scope
	rectangle := GetMinRectangle(originScope)

	fmt.Printf("MBR is %v \n", rectangle)

	geographicEngine.MBRScope = rectangle

	// 2.递归生成
	// 3.递归时候,应该要计算生成的网格是否在多边形内,以及他们的关系
	GenerateGridList(rectangle.MaxLat, rectangle.MinLng, rectangle.MaxLng, rectangle.MinLat, GEO_HASH_LEVEL_6)

	result := GenerateJSString(geographicEngine)
	fmt.Println(result)
}

// 递归生成网格列表
func GenerateGridList(lat, lng float64, maxLng float64, maxLat float64, level int64) {
	if lat < maxLat {
		return
	}
	// 1.这里的操作是将传入的初始值,计算成bounding box
	originGeoHash := geohash.Encode(lat, lng)
	box := geohash.BoundingBox(originGeoHash[:level])
	PolygonContains(box, true)
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

// 横向结构递归执行
func recursion(lat, lng float64, maxLng float64, level int64) {
	if lng > maxLng {
		return
	}
	boundingBox := ProduceBoundingBox(lat, lng, EAST, level)
	PolygonContains(boundingBox, true)
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

/*
1.多边形与多边形相交
2.多边形包含多边形
	a)生成的矩形四个点都在多边形内部

*/

// 校验多边形关系
func PolygonRelationship(Rectangle [][]float64) bool {
	// 如果四个点都不在的话,则排除
	flag := 0
	for _, point := range Rectangle {
		if !InPolygon(point, geographicEngine.OriginScope) {
			flag ++
		}
	}

	if flag == 4 {
		return false
	}

	return true
}

// 多边形包含
func PolygonContains(box geohash.Box, isOpen bool) {
	rectangle := generate(box)
	// 判断矩形的四个点是否都在MBR里面
	flag := PolygonRelationship(rectangle)
	if !flag {
		return
	}

	gridInfo := GridInfo{}
	if !isOpen {
		gridInfo.Scope = rectangle
		geographicEngine.GridList = append(geographicEngine.GridList, gridInfo)
		return
	}

	// 把第一个值塞到array里面
	rectangle = append(rectangle, rectangle[0])
	// 检查是否有交点,如果有交点求出改点构成的多边形
	result := checkPointInRectangle(geographicEngine.OriginScope, rectangle)
	//str := generatesJs(rectangle)
	if len(result) > 0 {
		gridInfo = GridInfo{
			Scope: result,
		}
	} else {
		//gridInfo = GridInfo{
		//	Scope: rectangle,
		//}
		return
	}

	geographicEngine.GridList = append(geographicEngine.GridList, gridInfo)

}

// 检查内
func checkPointInRectangle(originScope [][]float64, rectangle [][]float64) [][]float64 {
	// 1.先判断是否有交点
	// 2.有的话看点是否在圈内
	// 3.在圈内的点和交点就是这个多边形
	polygon := [][]float64{}
	pointList := CheckIntersection(originScope, rectangle)
	if len(pointList) > 0 {
		// 判断是否是顶点
		times := 0
		for i := 0; i < len(originScope)-1; i++ {
			if flag := InPolygon(originScope[i], rectangle); flag {
				times ++
				fmt.Println(rectangle)
				fmt.Println(flag,len(pointList))
			}
		}
		// 只有一个交点,并且有个点在四边形内
		if len(pointList) == 1 && times == 1 {
			fmt.Println("fuck")
		}

		polygon = append(polygon, pointList...)
		// 多边形的四个点是否在矩形内,最后一个点不要判断
		for i := 0; i < len(rectangle)-1; i++ {
			if flag := InPolygon(rectangle[i], originScope); flag {
				// 如果命中的话
				polygon = append(polygon, rectangle[i])
			}
		}

	}
	polygon = ClockwiseSortPoints(polygon)
	return polygon
}

//
func CheckIntersection(originScope [][]float64, rectangle [][]float64) [][]float64 {
	pointList := [][]float64{}
	for i := 0; i < len(originScope)-1; i++ {
		j := i + 1

		lineFirstStart := Point{
			X: originScope[i][0],
			Y: originScope[i][1],
		}
		lineFirstEnd := Point{
			X: originScope[j][0],
			Y: originScope[j][1],
		}
		for ii := 0; ii < len(rectangle)-1; ii++ {
			jj := ii + 1
			lineSecondStart := Point{
				X: rectangle[ii][0],
				Y: rectangle[ii][1],
			}
			lineSecondEnd := Point{
				X: rectangle[jj][0],
				Y: rectangle[jj][1],
			}

			result, err := GetIntersectionPoint(lineFirstStart, lineFirstEnd, lineSecondStart, lineSecondEnd)

			if err != nil {
				continue
			}
			// 校验命中的点是否合规
			if result == nil {
				continue
			}
			point, _ := checkPointRange(*result, rectangle)
			if point != nil {
				pointList = append(pointList, []float64{point.X, point.Y})
			}

		}
	}
	return pointList

}

// 检查交点是否在
func checkPointRange(point Point, rectangle [][]float64) (*Point, error) {

	// 116.39401012382888 39.990234375
	// [116.4111328125 39.9957275390625]
	// [116.4111328125 39.990234375]
	// [116.400146484375 39.990234375]
	// [116.38916015625 39.9957275390625
	r := GetMinRectangle(rectangle)
	if point.X > (r.MaxLng) || point.X < (r.MinLng) || point.Y > (r.MaxLat) || point.Y < (r.MinLat) {
		//fmt.Println(point, r, point.Y < (r.MinLat))
		return nil, errors.New("error happen")
	}

	return &point, nil

}

/*
* 直线方程L1: ( y - y1 ) / ( y2 - y1 ) = ( x - x1 ) / ( x2 - x1 )
             * => y = [ ( y2 - y1 ) / ( x2 - x1 ) ]( x - x1 ) + y1
             * 令 a = ( y2 - y1 ) / ( x2 - x1 )


*/
func GetIntersectionPoint(LineFirstStart Point, LineFirstEnd Point, LineSecondStart Point, LineSecondEnd Point) (*Point, error) {
	a := (LineFirstEnd.Y - LineFirstStart.Y) / (LineFirstEnd.X - LineFirstStart.X)
	b := Decimal(LineSecondEnd.Y-LineSecondStart.Y) / Decimal(LineSecondEnd.X-LineSecondStart.X)
	//fmt.Println(LineFirstStart, LineFirstEnd, LineSecondStart, LineSecondEnd, a, b)
	// 排除 + 、- Inf
	//if flag := math.IsInf(b, 1); flag {
	//	return &Point{}, errors.New("error")
	//}

	point := Point{}
	// math.Abs(b) == 0
	if math.IsInf(b, 0) {
		// b的斜率为0
		x := LineSecondStart.X;
		y := (LineFirstStart.X-x)*(-a) + LineFirstStart.Y
		point = Point{
			X: x,
			Y: y,
		}
		return &point, nil

	}

	x := (a*LineFirstStart.X - b*LineSecondStart.X - LineFirstStart.Y + LineSecondStart.Y) / (a - b)

	y := a*x - a*LineFirstStart.X + LineFirstStart.Y

	point = Point{
		X: Decimal(x),
		Y: Decimal(y),
	}
	return &point, nil
}

// 取小数点后几位
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.12f", value), 64)
	return value
}

func sortV2(rectangle [][]float64) [][]float64 {
	if len(rectangle) != 4 {
		return [][]float64{}
	}

	xAverage, yAverage, xSum, ySum := float64(0), float64(0), float64(0), float64(0)
	newRectangle := [][]float64{}

	for _, value := range rectangle {
		xSum += value[0]
		ySum += value[1]
	}
	xAverage = xSum / 4
	yAverage = ySum / 4

	for _, value := range rectangle {
		if value[0] < xAverage && value[1] < yAverage {
			newRectangle = append(newRectangle, value)
		}

	}
	for _, value := range rectangle {
		if value[0] > xAverage && value[1] < yAverage {
			newRectangle = append(newRectangle, value)
		}
	}
	for _, value := range rectangle {
		if value[0] > xAverage && value[1] > yAverage {
			newRectangle = append(newRectangle, value)
		}
	}
	for _, value := range rectangle {
		if value[0] < xAverage && value[1] > yAverage {
			newRectangle = append(newRectangle, value)
		}
	}
	return newRectangle
}

/*


 public PointF[] SortPointsByClockwise(PointF[] pts_src)
        {
            if (pts_src.Length != 4) return null;//确保为四边形

            //求四边形中心点？坐标
            float x_average = 0;
            float y_average = 0;
            float x_sum = 0;
            float y_sum = 0;
            for (int i = 0; i < 4; i++)
            {
                x_sum += pts_src[i].X;
                y_sum += pts_src[i].Y;
            }
            x_average = x_sum / 4;
            y_average = y_sum / 4;
            PointF center = new PointF(x_average, y_average);

            PointF[] result = new PointF[4];
            for (int i = 0; i < 4; i++)
            {
                if (pts_src[i].X < center.X && pts_src[i].Y < center.Y)
                {
                    result[0] = pts_src[i];//左上角点
                }
                if (pts_src[i].X > center.X && pts_src[i].Y < center.Y)
                {
                    result[1] = pts_src[i];//右上角点
                }
                if (pts_src[i].X > center.X && pts_src[i].Y > center.Y)
                {
                    result[2] = pts_src[i];//右下角点
                }
                if (pts_src[i].X < center.X && pts_src[i].Y > center.Y)
                {
                    result[3] = pts_src[i];//左下角点
                }
            }

            return result;



*/
