// surfaceは3-D面の関数のSVGレンダリングを計算します
package main

import (
	"fmt"
	"io"
	"math"
	"os"
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
var targetFigure string

type Point = struct {
	x float64
	y float64
}

func createSvg(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
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
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				a.x, a.y, b.x, b.y, c.x, c.y, d.x, d.y)
		}
	}
	fmt.Fprintln(w, "</svg>")
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
	var result float64
	switch targetFigure {
	case "saddle":
		result = saddle(x, y)
	case "eggbox":
		result = eggbox(x, y)
	default:
		result = sinrSurface(x, y)
	}
	if math.IsNaN(result) {
		return 0, false
	}
	return result, true
}

func sinrSurface(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func eggbox(x, y float64) float64 {
	return 0.1 * (math.Cos(x) + math.Sin(y))
}

func saddle(x, y float64) float64 {
	a := 25.0 * 25.0
	b := 17.0 * 17.0
	return y*y/a - x*x/b
}

func main() {
	var svgFileName string
	if len(os.Args) > 1 {
		targetFigure = os.Args[1]
	}
	switch targetFigure {
	case "saddle":
		svgFileName = "saddle.svg"
	case "eggbox":
		svgFileName = "eggbox.svg"
	default:
		svgFileName = "out.svg"
	}
	svgFile, err := os.Create(svgFileName)
	if err != nil {
		os.Exit(1)
	}
	defer svgFile.Close()

	createSvg(svgFile)
}
