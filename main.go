package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"go.k6.io/k6/cmd"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	cmd.Execute()
}
