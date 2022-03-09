package main

import (
	"fmt"
	"log"

	"github.com/sataapon/btc/internal/config"
	wlmigrate "github.com/sataapon/btc/internal/migrate"
)

func main() {
	_ = config.GetDB()

	err := wlmigrate.New().CreateTable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create table success")
}
