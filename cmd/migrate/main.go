package main

import (
	"fmt"

	wlmigrate "github.com/sataapon/btc/internal/migrate"
)

func main() {
	err := wlmigrate.New().CreateTable()
	if err != nil {
		panic(err)
	}
	fmt.Println("create table success")
}
