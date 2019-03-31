package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	sleepAndTalk(ctx, 5*time.Second, "hello")
}

func sleepAndTalk(ctx context.Contetx, t time.Duration, s String) {

}