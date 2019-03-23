package main

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

/** 
* Created by wanjx in 2019/3/21 10:04
**/
 
func TestWithTimeout(t *testing.T) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 100*time.Microsecond)

	req, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	fmt.Println("Response received, status code:", res.StatusCode)
}