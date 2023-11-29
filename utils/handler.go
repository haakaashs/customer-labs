package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GoHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	funcDesc := "GoHandler"
	log.Println("entered " + funcDesc)

	var req map[string]string

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("error while decode")
	}

	log.Println("request: ", req)
	channel := make(chan Responce)
	go goWorker(req, channel)
	responseData, _ := json.Marshal(<-channel)
	log.Println("exit " + funcDesc)

	// write response
	// w.WriteHeader(http.StatusOK)
	// w.Write(responseData)
	log.Println("response: ", responseData)
	webHooks("", responseData)
	log.Println("exit " + funcDesc)
}

func webHooks(webhookURL string, payload []byte) {
	funcDesc := "webHooks"
	log.Println("entered " + funcDesc)

	response, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}

	fmt.Println("Response Code:", response.Status)
	log.Println("exit " + funcDesc)
}
