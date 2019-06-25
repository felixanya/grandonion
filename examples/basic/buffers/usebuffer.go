package main

import (
	"bytes"
	"fmt"
)

/*
 Go中可以使用“+”合并字符串，但是这种合并方式效率非常低，每合并一次，都是创建一个新的字符串,就必须遍历复制一次字符串。
 使用bytes.Buffer来组装字符串，不需要复制，只需要将添加的字符串放在缓存末尾即可。
 Buffer是非线程安全的
*/

func main() {
	var buffer bytes.Buffer
	for i := 0; i < 100; i++ {
		buffer.WriteString("a")
	}

	fmt.Println(buffer.String())
}
