package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
函数 rand.Float32 和 rand.Float64 返回介于 [0.0, 1.0) 之间的伪随机数，其中包括 0.0 但不包括 1.0。函数 rand.Intn 返回介于 [0, n) 之间的伪随机数。

你可以使用 Seed(value) 函数来提供伪随机数的生成种子，一般情况下都会使用当前时间的纳秒级数字
*/

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	for i := 0; i < 10; i++ {
		r := rand.Intn(20)
		fmt.Printf("%d / ", r)
	}
	fmt.Println()

	times := int64(time.Now().Nanosecond())
	rand.Seed(times)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
}
