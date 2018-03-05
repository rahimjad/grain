package main

import (
	"fmt"

	"./api"
	"./db"
)

func main() {
	fmt.Println("Hello World")
	db.Init()
	api.Run()
}
