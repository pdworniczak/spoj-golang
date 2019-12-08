package main

import (
	"fmt"
	"strings"
)

func main() {
	var tests, n, x, y int

	fmt.Scan(&tests)
	results := make([][]int, tests)

	for i := 0; i < tests; i++ {
		fmt.Scan(&n, &x, &y)

		for r := 1; r < n; r++ {
			if r%x == 0 && r%y != 0 {
				results[i] = append(results[i], r)
			}
		}
	}

	for _, v := range results {
		fmt.Println(strings.Trim(fmt.Sprint(v), "[]"))
	}
}
