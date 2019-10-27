package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello,wold"
	fmt.Println(len(s))     // 10
	fmt.Println(s[0], s[7]) // 104 111

	f := "hello,世界"
	fmt.Println(len(f))     // 12
	fmt.Println(f[0], f[7]) // 104 184

	fmt.Println("substring ---------- ")
	fmt.Println(s[0:5]) // hello
	fmt.Println(s[:5])  //hello
	fmt.Println(s[5:])  //,wold

	fmt.Println("my" + s[5:]) // my,wold
	fmt.Println("escape ------ ")
	fmt.Println("\a,\b,\f,\n,\r,\t,\v,\",\\")

	fmt.Println("\xaf,\377")
	
	fmt.Println("utf-8 ---------")

	fmt.Println("世界")
	fmt.Println("\xe4\xb8\x96")
	fmt.Println("\u4e16\u754c")
	fmt.Println("\U00004e16\U0000754c")

	fmt.Println(hasPrefix("世界", "世"))

	decode("你好")
	decodeRange("世界")

	t := "你吃了吗？"
	r := []rune(t)
	fmt.Println(r, string(r))
	fmt.Printf("% x", t)


}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func decode(s string)  {
	for i := 0; i < len(s) ;{
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", r, r)
		i += size
	}
}

func decodeRange(s string) {
	for i, r := range s {
		fmt.Printf("%d\t%c\t%d\n", r, r, i)
	}
}
