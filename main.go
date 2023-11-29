package main

import (
	"customer-lobs/utils"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/post-go-worker", utils.GoHandler)
	go func() {
		// if any error while start simulation
		time.Sleep(time.Second * 2)
		log.Println("server started\nserving localhost port 4321")
	}()

	if err := http.ListenAndServe(":4321", nil); err != nil {
		log.Fatal("unable to start server")
	}
}
