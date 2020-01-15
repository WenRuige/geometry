## Golang-基于空间索引的地理围栏判定优化实现


> 背景



> 什么是地理围栏?

地理围栏（Geo-fencing）是LBS的一种新应用，就是用一个虚拟的栅栏围出一个虚拟地理边界


> 如何生成判定点是否在多边形内?


射线法 raw



> 如何快速匹配?
* geohash
* uberH3
* Google S2




> 最小外接矩形

最小外接矩形(minimum bounding rectangle,MBR),译为最小边界矩形
MBR

> 多边形边数较多怎么办?



> 判断线段是否在多边形内
* 首先，要判断一条线段是否在多边形内，先要判断线段的两个端点是否在多边形内。如果两个端点不全在多边形内，那么，线段肯定是不在多边形内的。
* 其次，如果线段和多边形的某条边内交（两线段内交是指两线段相交且交点不在两线段的端点），则线段肯定不在多边形内。
* 如果多边形的某个顶点和线段相交，则必须判断两相交交点之间的线段是否包含于多边形内。


> 线段是否与线段相交

跨立实验,

设矢量P=(x1,y1),Q=(x2,y2),则矢量叉积:PxQ = x1*y2 - x2*y1

叉积的一个非常重要性质是可以通过它的符号判断两矢量相互之间的顺逆时针关系
* 若PxQ > 0
* 若PxQ < 0
* 若PxQ = 0 



> 矢量叉积




> 数据库层的优化

mysql存储  

redis存储 / redis 分片存储


localcache存储



> 思考


### Reference
* https://mp.weixin.qq.com/s/2B-VJ2xgwxrmsSkE6zuoPA[闲鱼基于快速GeoHash实现海量商品与商圈高效匹配的算法]
* https://en.wikipedia.org/wiki/Geohash
* https://en.wikipedia.org/wiki/Pointinpolygon
* https://www.geeksforgeeks.org/how-to-check-if-a-given-point-lies-inside-a-polygon [4]https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-geohashgrid-aggregation.html
* http://blog.notdot.net/2009/11/Damn-Cool-Algorithms-Spatial-indexing-with-Quadtrees-and-Hilbert-Curves
