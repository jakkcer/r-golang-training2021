// surfaceは3-D面の関数のSVGレンダリングを計算します
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // キャンバスの大きさ (画素数)
	cells         = 100                 // 格子のます目の数
	xyrange       = 30.0                // 軸の範囲 (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x単位 および y単位当たりの画素数
	zscale        = height * 0.4        // z単位当たりの画素数
	angle         = math.Pi / 6         // x, y軸の角度 (=30度)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30度), cos(30度)

type Point = struct {
	x float64
	y float64
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	cornerMap := getCornerMap()

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			keyA := fmt.Sprintf("%d,%d", i+1, j)
			keyB := fmt.Sprintf("%d,%d", i, j)
			keyC := fmt.Sprintf("%d,%d", i, j+1)
			keyD := fmt.Sprintf("%d,%d", i+1, j+1)
			a, okA := cornerMap[keyA]
			b, okB := cornerMap[keyB]
			c, okC := cornerMap[keyC]
			d, okD := cornerMap[keyD]
			if !okA || !okB || !okC || !okD {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				a.x, a.y, b.x, b.y, c.x, c.y, d.x, d.y)
		}
	}
	fmt.Println("</svg>")
}

func getCornerMap() map[string]Point {
	cornerMap := make(map[string]Point, int(math.Pow(cells, 2)))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			if x, y, ok := corner(i, j); ok {
				cornerXY := fmt.Sprintf("%d,%d", i, j)
				cornerMap[cornerXY] = Point{x, y}
			}
		}
	}
	return cornerMap
}

func corner(i, j int) (float64, float64, bool) {
	// ます目(i,j)のかどの点(x,y)を見つける
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する
	z, ok := f(x, y)
	if !ok {
		return 0, 0, false
	}

	// (x,y,z)を2-D SVGキャンパス (sx, sy) へ等角的に投影
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y)
	result := math.Sin(r) / r
	if math.IsNaN(result) {
		return 0, false
	}
	return result, true
}
