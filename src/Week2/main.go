package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

/** 
* Created by wanjx in 2019/3/20 22:48
**/


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fmt.Fprint(os.Stdout, "处理请求：\n")

		select {
		case <- time.After(2 * time.Second):
			w.Write([]byte("请求处理完毕。"))
		case <-ctx.Done():
			fmt.Fprint(os.Stderr, "请求取消\n")
		}
	})
	http.ListenAndServe(":8000", nil)
}
