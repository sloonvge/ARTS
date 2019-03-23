package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

/** 
* Created by wanjx in 2019/3/21 9:46
**/
 
func operation1(ctx context.Context) error{
	time.Sleep(100 * time.Microsecond)
	return errors.New("failed")
}

func operation2(ctx context.Context){
	select {
	case <- time.After(500 * time.Microsecond):
		fmt.Println("done")
	case <- ctx.Done():
		fmt.Println("halted operation2")
		
	}
}

func main(){
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	var wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		err := operation1(ctx)
		if err != nil {
			cancel()
		}
		wg.Done()
	}()

	operation2(ctx)
}