package main

/*
   * @description 射线法判断点是否在多边形内部
   * @param {Object} p 待判断的点，格式：{ x: X坐标, y: Y坐标 }
   * @param {Array} poly 多边形顶点，数组成员的格式同 p
   * @return {String} 点 p 和多边形 poly 的几何关系
function rayCasting(p, poly) {
var px = p.x,
py = p.y,
flag = false

for(var i = 0, l = poly.length, j = l - 1; i < l; j = i, i++) {
var sx = poly[i].x,
sy = poly[i].y,
tx = poly[j].x,
ty = poly[j].y

// 点与多边形顶点重合
if((sx === px && sy === py) || (tx === px && ty === py)) {
return 'on'
      }

// 判断线段两端点是否在射线两侧
if((sy < py && ty >= py) || (sy >= py && ty < py)) {
// 线段上与射线 Y 坐标相同的点的 X 坐标
var x = sx + (py - sy) * (tx - sx) / (ty - sy)

// 点在多边形的边上
if(x === px) {
return 'on'
        }

// 射线穿过多边形的边界
if(x > px) {
flag = !flag
}
}
}

// 射线穿过多边形边界的次数为奇数时点在多边形内
return flag ? 'in' : 'out'
  }

*/

func euqal2(x float64, y float64) bool {
	v := x - y
	const delta float64 = 1e-6
	if v < delta && v > -delta {
		return true
	}
	return false

}

// 射线法
func rayCasting(point []float64, vertices [][]float64) string {
	px := point[0]
	py := point[1]
	flag := false

	if flag {

	}

	l := len(vertices)
	for i := 0; i < l; i++ {
		j := l - 1

		sx := vertices[i][0]
		sy := vertices[i][1]

		tx := vertices[j][0]
		ty := vertices[j][1]

		j=i
		// 如果点在多边形上面
		if (sx == px && sy == py) || (tx == px && ty == py) {
			return "on"
		}
		// 判断两个端点是否在射线的两侧
		if ((sy < py && ty >= py) || (sy >= py && ty < py)) {
			flag = true

			x := sx + (py-sy)*(tx-sx)/(ty-sy)

			if x == px {
				return "on"
			}

			if x > px {
				flag = !flag
			}
		}



	}
	return "off"
}

/*
点是否在多边形上实现思路
1.点是否落在多边形顶点上了

*/

func main() {

}

// (sx,xy),(tx,ty)
// y = kx + b

// 斜截式 y = kx +b  其中k是直线的斜率,b是直线在y轴上的截距.该方程叫做直线的斜截式方程,简称斜截式

// sy = sx * k +b
// ty = tx * k +b

// => sy - ty = k(sx-tx)
// => k = (sx-tx)/(sy-ty)

// html-js.com/article/1528
