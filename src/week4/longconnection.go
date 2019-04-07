package main

import (
	"net"
	"time"
	"fmt"
	"net/http"
	"io/ioutil"
)

const URL = "http://localhost:8888/"

func LocalDial(network, addr string) (net.Conn, error) {
	dial := net.Dialer{
		Timeout: 30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	conn, err := dial.Dial(network, addr)
	if err != nil {
		return conn, err
	}

	fmt.Println("connect done, use", conn.LocalAddr().String())

	return conn, err
}

func doGet(client *http.Client, url string, id int) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%d:%s -- %v\n", id, string(buf), err)
	if err := resp.Body.Close(); err != nil {
		fmt.Println(err)
	}
}


func doGetFailLong(client *http.Client, url string, id int) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// buf, err := ioutil.ReadAll(resp.Body)

	if err := resp.Body.Close(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d:%s\n", id, "done")
}

func main() {

	client := &http.Client{
		Transport: &http.Transport{
			Dial: LocalDial,
		},
	}
	for {
		go doGetFailLong(client, URL, 1)
		go doGetFailLong(client, URL, 2)
		go doGetFailLong(client, URL, 3)
		time.Sleep(2 * time.Second)
	}
}