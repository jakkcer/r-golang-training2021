package main

import "fmt"

func noReturn() (err error) {
	type fakePanic struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// パニックなし
		case fakePanic{}:
			// "no return"を返す
			err = fmt.Errorf("no return")
		default:
			panic(p) // 予期しないパニック; パニックを維持する
		}
	}()

	panic(fakePanic{})
}

func main() {
	fmt.Println(noReturn())
}
