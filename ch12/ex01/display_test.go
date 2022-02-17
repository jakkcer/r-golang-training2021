package display

func displayStructMap() {
	structMap := map[struct{ x int }]int{
		{1}: 2,
		{2}: 3,
	}
	Display("structMap", structMap)
}

func Example_displayStructMap() {
	displayStructMap()
	// Output:
	// Display structMap (map[struct { x int }]int):
	// structMap[{x: 1}] = 2
	// structMap[{x: 2}] = 3
}

func displayArrayMap() {
	arrayMap := map[[3]int]int{
		{1, 2, 3}: 3,
		{2, 3, 4}: 4,
	}
	Display("arrayMap", arrayMap)
}

func Example_displayArrayMap() {
	displayArrayMap()
	// Output:
	// Display arrayMap (map[[3]int]int):
	// arrayMap[{1, 2, 3}] = 3
	// arrayMap[{2, 3, 4}] = 4
}
