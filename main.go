package main

import (
	"fmt"
	"math"
	"monitor/util"
	"runtime"
	"time"
)

type user struct {
	name string
	age  int
}

func main() {
	fmt.Printf("guo\n")
	var i bool

	d := 1
	fmt.Println(i)
	fmt.Println(Tes(10, d))
	ui := float64(d)
	math.Sqrt(float64(ui))
	fmt.Println(d)
	if x := 1; !i {
		fmt.Printf("1111")
		fmt.Println(x)
	}
	fmt.Println(d)

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	var uu user = user{name: "gzf", age: d}
	fmt.Println(uu.name)
	fmt.Println(reverse(10, 11))

	time.After(2 * time.Second)
	fmt.Println(11)
	post, str := util.HttpPost("http://www.baidu.com", "111")
	fmt.Println(string(post))
	fmt.Println(str)
}

func Tes(x, y int) int {
	return x + y
}
func reverse(x, y int) (int, int) {
	return y, x
}

func init() {
	fmt.Printf("fffffffffffffffffffffffff")
	var d user = user{"11", 10}
	fmt.Println(d.name)
}
