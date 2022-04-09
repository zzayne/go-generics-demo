package main

import "fmt"

type num interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func sum[T num](arr []T) T {
	var r T
	for _, v := range arr {
		r += v
	}
	return r
}

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(sum([]float64{1.0, 3.4, 5.6}))
}
