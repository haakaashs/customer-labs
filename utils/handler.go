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

	log.Println("response: ", string(responseData))
	webHooks("https://webhook.site/4ad1a5cd-97a0-4226-b628-3ece50ba7cc5", responseData)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"OK","message":"success"}`))
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

	log.Println("weebHooks Response Code:", response.Status)
	log.Println("exit " + funcDesc)
}
