package witbot

import (
	"github.com/claudioontheweb/witbot/config"
	witai "github.com/wit-ai/wit-go"
	"net/http"
)

type Witbot struct {
	Config 		*config.Config
	HttpClient	*http.Client
	WitaiClient	*witai.Client
}

type WitResponse struct {
	Text		string`json:"_text"`
	Entities	map[string]interface{}`json:"entities"`
	MsgId		string`json:"msg_id"`
}


func NewWitbot(config config.Config) (*Witbot, error) {

	// Configure HTTP Client from config
	httpClient := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	client := witai.NewClient(config.WIT_AI_TOKEN)
	client.SetHTTPClient(httpClient)

	witbot := &Witbot{
		Config:      &config,
		HttpClient:  httpClient,
		WitaiClient: client,
	}

	return witbot, nil
}

// Checks what to do with the Witbot AI HTTP Response
func handleWitResponse(response WitResponse) (interface{}, error) {

	return response, nil

}