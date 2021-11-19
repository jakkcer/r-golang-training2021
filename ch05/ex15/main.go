package number

import (
	"fmt"
)

func max(vals ...int) int {
	var maxNum int
	for i, val := range vals {
		if i == 0 {
			maxNum = val
		}
		if val > maxNum {
			maxNum = val
		}
	}
	return maxNum
}

func min(vals ...int) int {
	var minNum int
	for i, val := range vals {
		if i == 0 {
			minNum = val
		}
		if val < minNum {
			minNum = val
		}
	}
	return minNum
}

func argsRequiredMax(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("[Error] Need to pass at least one argument")
	}
	return max(vals...), nil
}

func argsRequiredMin(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("[Error] Need to pass at least one argument")
	}
	return min(vals...), nil
}
