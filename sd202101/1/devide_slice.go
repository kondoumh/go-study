package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	size := 2
	dst := make([][]int, 0, (len(src)+size-1)/size)
	for size < len(src) {
		src, dst = src[size:], append(dst, src[0:size:size])
	}
	dst = append(dst, src)
	fmt.Println(dst)
	fmt.Println(dst[0])
}
