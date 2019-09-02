package main

import (
	"fmt"
)

// CanTakeCakes 円状に並んだ16個のケーキを隣合わないように取れるかどうか
func CanTakeCakes(A int, B int) bool {
	// (A, B) = (1, 1), (1, 2),..., (1, 15)
	const N = 16
	if A > B {
		B, A = A, B
	}
	// A <= B として良い
	if B > int(N/2) {
		return false
	}
	var N2 = N - 2*A
	var B2 = B - A

	// N2個のケーキからB2個を取り分けられるか(端はB)
	if 2*B2 <= N2 {
		return true
	}
	return false
}

func main() {
	var A int
	var B int
	var canTake bool

	fmt.Scan(&A)
	fmt.Scan(&B)
	canTake = CanTakeCakes(A, B)
	if canTake {
		fmt.Println("Yay!")
	} else {
		fmt.Println(":(")
	}
}
