package main

import (
	"github.com/haakaashs/customer-labs/utils"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/go-worker", utils.GoHandler)
	go func() {
		// if any error while start simulation
		time.Sleep(time.Second * 2)
		log.Println("server started\nserving localhost port 4321")
	}()

	if err := http.ListenAndServe(":4321", nil); err != nil {
		log.Fatal("unable to start server")
	}
}
