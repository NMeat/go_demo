package demo

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTest1(t *testing.T) {
	// 返回一个空的 Context，这个空的 Context 一般用于整个 Context 树的根节点。
	ctx, cancel := context.WithCancel(context.Background())

	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				// 判断是否要结束，如果接受到值的话，就可以返回结束 goroutine 了
				case <-ctx.Done():
					fmt.Println("我要停止了")
					return
				case ch <- i:
					fmt.Printf("i = %d\n", i)
					break
				}
			}
		}()
		return ch
	}(ctx)

	for v := range ch {
		if v == 5 {
			// 发出取消指令
			cancel()
			break
		}
	}
	time.Sleep(time.Second * 4)
	fmt.Println("程序已要结束了")
}
