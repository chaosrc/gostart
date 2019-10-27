package main

import (
	"fmt"
)

func months() {
	months := [...]string{1: "January", 2: "February",
		3: "March", 4: "April", 5: "May",
		6: "June", 7: "July", 8: "August",
		9: "September", 10: "October", 11: "November",
		12: "December"}

	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Printf("%s len: %d, cap: %d\n", Q2, len(Q2), cap(Q2))
	fmt.Printf("%s len: %d, cap: %d\n", summer, len(summer), cap(summer))

	fmt.Println(Q2[:5])

	s := "hello world"
	b := []byte(s)
	s2 := b[2:4]
	fmt.Println(s2, len(s2), cap(s2))
}

func reverse(s []int) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
}
