// 重さをポンドとキログラムで変換します．
package weightconv

import "fmt"

type Pound float64
type Kilogram float64

func (p Pound) String() string     { return fmt.Sprintf("%glb", p) }
func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }

func PToKg(p Pound) Kilogram  { return Kilogram(p / 2.205) }
func KgToP(kg Kilogram) Pound { return Pound(kg * 2.205) }
