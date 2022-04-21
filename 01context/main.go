package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 如何通知 子goroutine 退出
// 1、使用全局变量bool
// 2、使用管道
// 3、使用 context

var waitGroup sync.WaitGroup

//var ch = make(chan bool, 1) // 使用管道
var flag bool // 使用全局变量bool

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background()) // 使用 ctx
	// 如何通知 子goroutine 退出
	waitGroup.Add(1)
	go simpleTask(ctx)
	time.Sleep(time.Second * 5)
	cancelFunc() // 通知取消函数执行
	waitGroup.Wait()
	fmt.Println("测试", flag)
	fmt.Println("测试正常合并")
	fmt.Println("master更改")
	fmt.Println("hot fix 修改")
}

// 一个简单的任务
func simpleTask(ctx context.Context) {

	defer waitGroup.Done()
	for {
		fmt.Println("why need context?")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
