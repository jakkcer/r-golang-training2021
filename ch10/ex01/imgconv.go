package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func main() {
	var format string
	flag.StringVar(&format, "f", "", "(required) 出力形式を指定してください．png, jpg, gifなど")
	flag.Parse()
	if len(flag.Args()) > 0 {
		fmt.Fprintln(os.Stderr, "usage: imgconv -f=png|jpg|gif < INPUT > OUTPUT")
		os.Exit(1)
	}

	img, _, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	format = strings.ToLower(format)
	switch format {
	case "jpg", "jpeg":
		err = jpeg.Encode(os.Stdout, img, nil)
	case "png":
		err = png.Encode(os.Stdout, img)
	case "gif":
		err = gif.Encode(os.Stdout, img, nil)
	default:
		err = fmt.Errorf("not supported")
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
