package witbot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func (w *Witbot) HandleReqs(res http.ResponseWriter, req *http.Request) {

	// Get Request Body and convert to string
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(req.Body)
	}
	// Restore the io.ReadCloser to its original state
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Use the content
	bodyString := string(bodyBytes)

	// Replace white spaces with %20
	// Replace line breaks
	bodyString = strings.ReplaceAll(bodyString, " ", "%20" )
	bodyString = strings.ReplaceAll(bodyString, "\n", "")

	// Prepare request to Wit.AI
	request, err := http.NewRequest("GET", "https://api.wit.ai/message?v=20191203&q=" + bodyString, nil)
	if err != nil {
		log.Fatal("err", err)
	}

	// Set Auth Header
	request.Header.Set("Authorization", "Bearer " + w.Config.WIT_AI_TOKEN)

	// Create HTTP Client
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	// Perform request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Get response Body and parse to WitResponse struct
	witbotRes := WitResponse{}
	err = json.NewDecoder(response.Body).Decode(&witbotRes)
	if err != nil {
		log.Fatal(err)
	}

	// Handle Wit Response and get result
	r, err := handleWitResponse(witbotRes)
	if err != nil {
		log.Fatal(err)
	}

	// Convert result to JSON and send back to client
	err = json.NewEncoder(res).Encode(r)
	if err != nil {
		panic(err)
	}
}