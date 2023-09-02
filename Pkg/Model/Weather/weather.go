package Weather

type WeatherData struct {
	Location Location       `json:"location"`
	Current  CurrentWeather `json:"current"`
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Latitude       float64 `json:"lat"`
	Longitude      float64 `json:"lon"`
	TimeZoneID     string  `json:"tz_id"`
	LocalTimeEpoch int64   `json:"localtime_epoch"`
	LocalTime      string  `json:"localtime"`
}

type CurrentWeather struct {
	LastUpdatedEpoch int64     `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TemperatureC     float64   `json:"temp_c"`
	TemperatureF     float64   `json:"temp_f"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMph          float64   `json:"wind_mph"`
	WindKph          float64   `json:"wind_kph"`
	WindDegree       int       `json:"wind_degree"`
	WindDirection    string    `json:"wind_dir"`
	PressureMB       float64   `json:"pressure_mb"`
	PressureIn       float64   `json:"pressure_in"`
	PrecipitationMm  float64   `json:"precip_mm"`
	PrecipitationIn  float64   `json:"precip_in"`
	Humidity         int       `json:"humidity"`
	Cloud            int       `json:"cloud"`
	FeelsLikeC       float64   `json:"feelslike_c"`
	FeelsLikeF       float64   `json:"feelslike_f"`
	VisibilityKm     float64   `json:"vis_km"`
	VisibilityMiles  float64   `json:"vis_miles"`
	UVIndex          float64   `json:"uv"`
	GustMph          float64   `json:"gust_mph"`
	GustKph          float64   `json:"gust_kph"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}
