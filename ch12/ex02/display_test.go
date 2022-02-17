package display

func ExampleDisplay() {
	// 循環しているstruct
	type Cycle struct {
		Value int
		Link  *Cycle
	}
	var c Cycle
	c = Cycle{16, &c}
	Display("c", c)
	// Output:
	// Display c (display.Cycle):
	// c.Value = 16
	// (*c.Link).Value = 16
	// (*(*c.Link).Link).Value = 16
	// (*(*(*c.Link).Link).Link) = display.Cycle value
}
