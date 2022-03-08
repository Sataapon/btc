package main

import (
	"fmt"
	"log"

	"github.com/sataapon/btc/internal/config"
	wlhttp "github.com/sataapon/btc/internal/wallet/http"
)

func main() {
	srv := wlhttp.NewServer(fmt.Sprintf("%s:%s", config.GetHostName(), config.GetServicePort()))
	fmt.Println("start service...")

	log.Fatal(srv.ListenAndServe())
}
