// 数値引数の単位変換を行います．
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"gobook/ch02/ex02/lengthconv"
	"gobook/ch02/ex02/weightconv"

	"gopl.io/ch2/tempconv"
)

var out io.Writer = os.Stdout

func main() {
	if len(os.Args[1:]) != 0 {
		for _, arg := range os.Args[1:] {
			convertAndShow(arg)
		}
	} else {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			convertAndShow(sc.Text())
		}
	}
}

func convertAndShow(arg string) {
	v, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
		os.Exit(1)
	}
	convTempature(v)
	convLength(v)
	convWeight(v)
}

func convTempature(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Fprintf(out, "%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func convLength(l float64) {
	f := lengthconv.Feet(l)
	m := lengthconv.Meter(l)
	fmt.Fprintf(out, "%s = %s, %s = %s\n", f, lengthconv.FToM(f), m, lengthconv.MToF(m))
}

func convWeight(w float64) {
	p := weightconv.Pound(w)
	kg := weightconv.Kilogram(w)
	fmt.Fprintf(out, "%s = %s, %s = %s\n", p, weightconv.PToKg(p), kg, weightconv.KgToP(kg))
}
