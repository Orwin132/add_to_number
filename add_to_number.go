package comp

import (
	"fmt"
	"strconv"
)

func add_to_number(n int) int {
	var resultat = n + 10

	num, err := strconv.ParseInt("11011", 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if n < int(num) {
		resultat = n
	}
	for i := 0; i <= n; i++ {
		resultat += i
	}
	return resultat
}
