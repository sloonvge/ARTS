package main

import (
	"net/http"
	"time"
	"testing"
)



func TestDoGetFailLong(t *testing.T) {

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