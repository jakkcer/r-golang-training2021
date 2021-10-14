// mandelbrotはマンデルブロフラクタルのフルカラーPNG画像を生成します
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	pngFile, err := os.Create("out.png")
	if err != nil {
		fmt.Printf("error on os.Create: %s", err)
		os.Exit(1)
	}
	defer pngFile.Close()

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点 (px, py) は複素数値zを表している
			img.Set(px, py, mandelbrot(z))

		}
	}
	if err := png.Encode(pngFile, img); err != nil {
		fmt.Printf("error on png.Encode: %s", err)
		os.Exit(1)
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 10

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			alpha := uint8(float64(n+1)/iterations*55.0) + 200
			nc := uint8((n % contrast) * alpha / contrast)
			switch n % 3 {
			case 0:
				return color.RGBA{nc, alpha, alpha, alpha}
			case 1:
				return color.RGBA{alpha, nc, alpha, alpha}
			case 2:
				return color.RGBA{alpha, alpha, nc, alpha}
			}
		}
	}
	return color.RGBA{0, 0, 0, 200}
}
