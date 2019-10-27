package main

import (
	"fmt"
)

func ap() {
	var s []rune

	for _, item := range "hello 世界" {
		s = append(s, item)
	}
	fmt.Printf("%q\n", s)
}

func appendInt(src []int, val int) []int {
	var dist []int

	zlen := len(src) + 1

	if zlen <= cap(src) {
		// 有增长空间，则扩展 slice
		dist = src[:zlen]
	} else {
		// 空间不足，创建新的数组，容量扩大一倍
		zcap := 2 * len(src)
		dist = make([]int, zlen, zcap)
		copy(dist, src)
	}
	dist[zlen-1] = val
	return dist
}

func appendMulti() {
	s := []int{1}
	s = append(s, 2, 3)
	s = append(s, 4, 5, 6)
	fmt.Println(s)
}

func noempty(str []string) []string {
	var i int

	for _, s := range str {
		if s != "" {
			str[i] = s
			i++
		}
	}
	return str[:i]
}

func noempty2(str []string) []string {
	var r []string

	for _, s := range str {
		if s != "" {
			r = append(r, s)
		}
	}
	return r
}
