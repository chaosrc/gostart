package echo

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "不换行")
var s = flag.String("s", " ", "分隔符")

// Start echo
func Start() {
	flag.Parse()
	text := strings.Join(flag.Args(), *s)
	if *n == true {
		fmt.Printf(text)
	} else {
		fmt.Println(text)
	}
}
