package weather

// Struct coming from Wit.ai
type WitWeather struct {
	Datetime	[]Datetime
	Location	[]Location
	Weather		[]Weather
}

type Datetime struct {

	Confidence	float64
	Grain		string
	Type		string
	Value		string
	Values		[]datetimeValues

}

type datetimeValues struct {
	Grain	string
	Type	string
	Value	string
}


type Location struct {
	Confidence	float64
	Resolved	struct{
		Values	[]locationValues
	}
	Value		string
}

type locationValues struct {
	Coords		struct {
		Lat			float64
		Long		float64
	}
	External	struct {
		Geonames	string
		Wikidata	string
		Wikipedia	string
	}
	Grain		string
	Name		string
	Timezone	string
	Type		string
}

type Weather struct {
	Confidence	float64
	Type		string
	Value		string
}

// Response to HTTP Client
type WeatherResponse struct {
	Location	string
	Date		string
}