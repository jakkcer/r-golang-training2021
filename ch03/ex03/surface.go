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

var (
	sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30度), cos(30度)
	zmin, zmax   = math.Inf(1), math.Inf(-1)
	cornerMap    = make(map[string]Point, int(math.Pow(cells, 2)))
	colorMap     = make(map[string]string, int(math.Pow(cells, 2)))
)

type Point = struct {
	x      float64
	y      float64
	height float64
}

func createSvg(w io.Writer) {
	calcCornerMap()

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

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
			color := getColor(a, b, c, d)
			fmt.Fprintf(w, "<polygon fill='%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, a.x, a.y, b.x, b.y, c.x, c.y, d.x, d.y)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func calcCornerMap() {
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			if point, ok := corner(i, j); ok {
				cornerXY := fmt.Sprintf("%d,%d", i, j)
				cornerMap[cornerXY] = point
			}
		}
	}
}

func getColor(a, b, c, d Point) string {
	// 4点の最大，最小をとる
	min, max := math.Inf(1), math.Inf(-1)
	for _, p := range []Point{a, b, c, d} {
		if p.height < min {
			min = p.height
		}
		if p.height > max {
			max = p.height
		}
	}
	var red, blue float64
	if max >= 0 {
		red = (max/zmax + 1) * 127.5
		blue = 255.0 - red
	} else {
		blue = (min/zmin + 1) * 127.5
		red = 255.0 - blue
	}
	return fmt.Sprintf("#%02x00%02x", uint8(red), uint8(blue))
}

func corner(i, j int) (Point, bool) {
	// ます目(i,j)のかどの点(x,y)を見つける
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する
	z, ok := f(x, y)
	if !ok {
		return Point{}, false
	}
	if z < zmin {
		zmin = z
	}
	if z > zmax {
		zmax = z
	}

	// (x,y,z)を2-D SVGキャンパス (sx, sy) へ等角的に投影
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return Point{sx, sy, z}, true
}

func f(x, y float64) (float64, bool) {
	result := sinrSurface(x, y)
	if math.IsNaN(result) {
		return 0, false
	}
	return result, true
}

func sinrSurface(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func main() {
	svgFile, err := os.Create("out.svg")
	if err != nil {
		os.Exit(1)
	}
	defer svgFile.Close()

	createSvg(svgFile)
}
