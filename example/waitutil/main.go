package main

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func main() {
	ctx := context.Background()

	go wait.UntilWithContext(ctx, test1, time.Minute)
	wait.UntilWithContext(ctx, test2, 0)
}

func test1(ctx context.Context) {
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(t + " test1")
	time.Sleep(5 * time.Minute)
	fmt.Println(t + "test1 end")
}

func test2(ctx context.Context) {
}
