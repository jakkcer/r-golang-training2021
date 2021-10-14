// serverは3-D面の関数のSVGレンダリングを計算してSVGを書き出すサーバです
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // 格子のます目の数
	xyrange = 30.0        // 軸の範囲 (-xyrange..+xyrange)
	angle   = math.Pi / 6 // x, y軸の角度 (=30度)
)

var (
	xyscale       float64                            // x単位 および y単位当たりの画素数
	zscale        float64                            // z単位当たりの画素数
	width, height = 600.0, 320.0                     // キャンバスの大きさ (画素数)
	sin30, cos30  = math.Sin(angle), math.Cos(angle) // sin(30度), cos(30度)
	fillColor     = "white"                          // svgの色
	zmin, zmax    = math.Inf(1), math.Inf(-1)
	cornerMap     = make(map[string]Point, int(math.Pow(cells, 2)))
)

type Point = struct {
	x float64
	y float64
}

func main() {
	http.HandleFunc("/", createSVGHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func createSVGHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	if qWidth := r.URL.Query().Get("width"); qWidth != "" {
		if v, err := strconv.ParseFloat(qWidth, 64); err == nil {
			width = v
		}
	}
	if qHeight := r.URL.Query().Get("height"); qHeight != "" {
		if v, err := strconv.ParseFloat(qHeight, 64); err == nil {
			height = v
		}
	}
	if qColor := r.URL.Query().Get("color"); qColor != "" {
		fillColor = qColor
	}
	createSVG(w)
}

func createSVG(w io.Writer) {
	xyscale = width / 2 / xyrange
	zscale = height * 0.4
	calcCornerMap()

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", fillColor, int(width), int(height))

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

func corner(i, j int) (Point, bool) {
	// ます目(i,j)のかどの点(x,y)を見つける
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する
	z, ok := f(x, y)
	if !ok {
		return Point{}, false
	}
	// (x,y,z)を2-D SVGキャンパス (sx, sy) へ等角的に投影
	sx := width/2.0 + (x-y)*cos30*float64(xyscale)
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return Point{sx, sy}, true
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
