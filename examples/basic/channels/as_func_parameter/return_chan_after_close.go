package as_func_parameter

import "fmt"

func returnChan() (chan string, error) {
	ch := make(chan string, 10)
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("第%d个", i)
	}

	return ch, nil
}
