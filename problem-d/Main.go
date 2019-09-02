package main

import (
	"fmt"
	"math"
	"sort"
)

// solve (-1, 1)を事前にかけたものをソート
func solve(N int, M int, arr [][]int) int {
	var maxval int
	maxval = math.MinInt64
	for ix := -1; ix <= 1; ix += 2 {
		for iy := -1; iy <= 1; iy += 2 {
			for iz := -1; iz <= 1; iz += 2 {
				converted := make([]int, N)
				for i := 0; i < N; i++ {
					a := arr[i]
					converted[i] = a[0]*int(ix) + a[1]*int(iy) + a[2]*int(iz)
				}
				sort.Ints(converted)
				var val int
				val = 0
				for i := 0; i < M; i++ {
					val += converted[N-1-i]
				}
				if maxval < val {
					maxval = val
				}
			}
		}
	}
	return maxval
}

func main() {
	var (
		N int
		M int
	)

	fmt.Scan(&N)
	fmt.Scan(&M)
	arr := make([][]int, N)
	for i := 0; i < int(N); i++ {
		arr[i] = make([]int, 3)
		fmt.Scan(&arr[i][0])
		fmt.Scan(&arr[i][1])
		fmt.Scan(&arr[i][2])
	}
	res := solve(N, M, arr)
	fmt.Println(res)
}
