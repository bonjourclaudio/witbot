package witbot

import (
	"encoding/json"
	"github.com/claudioontheweb/witbot/config"
	"github.com/claudioontheweb/witbot/pkg/weather"
	witai "github.com/wit-ai/wit-go"
	"net/http"
	"reflect"
)

type Witbot struct {
	Config 		*config.Config
	HttpClient	*http.Client
	WitaiClient	*witai.Client
}

type WitResponse struct {
	MsgID		string`json:"msg_id"`
	Text		string`json:"_text"`
	Entities 	interface{}`json:"entities"`
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
func handleWitResponse(response *WitResponse) (interface{}, error) {

	// Get entities as map string interface
	entities := reflect.ValueOf(response.Entities).Interface().(map[string]interface{})


	// check if weather is being requested
	if _, ok := entities["weather"]; ok {

		// parse response to WitWeather
		jsonbody, err := json.Marshal(entities)
		if err != nil {
			panic(err)
		}

		res := weather.WitWeather{}

		if err := json.Unmarshal(jsonbody, &res); err != nil {
			panic(err)
		}

		// handle weather action
		return HandleWeather(&res)

	} else {
		response := struct {
			Message string
		}{
			"No match found...",
		}
		return response, nil
	}

	return response, nil

}

func HandleWeather(w *weather.WitWeather) (weather.WeatherResponse, error) {
	wr := weather.WeatherResponse{}
	if len(w.Location) > 0 {
		wr.Location = w.Location[0].Value
		wr.Date = w.Datetime[0].Value
	}

	return wr, nil
}