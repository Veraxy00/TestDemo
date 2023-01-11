package main

import "webDemo/router"

const (
	token = "iiiia.txt"
)

func main() {
	if err := router.Start(); err != nil {
		panic(err)
	}
}
