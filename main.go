package main

import "webDemo/router"

const (
	token = "a.txt"
)

func main() {
	if err := router.Start(); err != nil {
		panic(err)
	}
}
