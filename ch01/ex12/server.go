// serverはリサジュー図形を返すサーバーです
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

type LissajousParams struct {
	cycles  float64 // 発振器xが完了する周回の回数
	res     float64 // 回転の分解能
	size    int     // 画像キャンバスは[-size..+size]の範囲を扱う
	nframes int     // アニメーションフレーム数
	delay   int     // 10ms単位でのフレーム間の遅延
}

var (
	palette    = []color.Color{color.Black, colorGreen}
	colorGreen = color.RGBA{0x00, 0xff, 0x00, 0xff}
)

const (
	backgroundIndex = 0 // パレットの背景の色
	lineIndex       = 1 // パレットの線の色
)

func main() {
	http.HandleFunc("/", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	params := LissajousParams{
		cycles:  5,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}
	for key, vs := range r.URL.Query() {
		switch key {
		case "cycles", "res":
			if v, err := strconv.ParseFloat(vs[0], 64); err == nil {
				if key == "cycles" {
					params.cycles = v
				} else if key == "res" {
					params.res = v
				}
			}
		case "size", "nframes", "delay":
			if v, err := strconv.Atoi(vs[0]); err == nil {
				if key == "size" {
					params.size = v
				} else if key == "nframes" {
					params.nframes = v
				} else if key == "delay" {
					params.delay = v
				}
			}
		}
	}
	fmt.Println(params)
	lissajous(w, params)
}

func lissajous(out io.Writer, params LissajousParams) {
	freq := rand.Float64() * 3.0 // 発振器yの相対周波数
	anim := gif.GIF{LoopCount: params.nframes}
	phase := 0.0 // 位相差

	for i := 0; i < params.nframes; i++ {
		rect := image.Rect(0, 0, 2*params.size+1, 2*params.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < params.cycles*2*math.Pi; t += params.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(params.size+int(x*float64(params.size)+0.5), params.size+int(y*float64(params.size)+0.5), lineIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, params.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意: エンコードエラーを無視
}
