package Weather

import (
	"Thalox/Pkg/Client"
	WeatherModel "Thalox/Pkg/Model/Weather"
	"Thalox/Pkg/Utils"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"net/url"
	"strconv"
)

func Weather(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	request, err := Client.CreateRequest(buildRequest(args[1]), "GET", nil, nil)
	if err != nil {
		return
	}

	response, err := Client.Response(request)
	if err != nil {
		return
	}

	var weatherData WeatherModel.WeatherData
	if err := json.Unmarshal(response, &weatherData); err != nil {
		return
	}
	weatherParse(weatherData)
	embed := &discordgo.MessageEmbed{
		Title: "<@" + m.Author.ID + "> " + weatherData.Location.Name + " Hava Durumu Bilgisi aşağıda yer almaktadır",
		Description: "İweatherData.Location.Region'da hava " + weatherData.Current.Condition.Text + " (" + strconv.FormatFloat(weatherData.Current.TemperatureC, 'f', -1, 64) + "°C) bir gün. Rüzgar hafifçe Kuzey Kuzeydoğu'dan esiyor. " +
			"Gökyüzünde neredeyse hiç bulut yok, görüş mesafesi 10 kilometre ve nem oranı yüzde 38. Dışarıda zaman geçirmek için ideal bir hava.",
		Color: 0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Bügün",
				Value:  fmt.Sprintf("%s °C", strconv.FormatFloat(weatherData.Current.TemperatureC, 'f', -1, 64)),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Yarın",
				Value:  fmt.Sprintf("%s °C", strconv.FormatFloat(weatherData.Current.TemperatureC, 'f', -1, 64)),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Sonraki gün",
				Value:  fmt.Sprintf("%s °C", strconv.FormatFloat(weatherData.Current.TemperatureC, 'f', -1, 64)),
				Inline: true,
			},
		},
		//Image: &discordgo.MessageEmbedImage{
		//	URL: "",
		//},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
func weatherParse(weatherData WeatherModel.WeatherData) {
	as := weatherData
	fmt.Println(as)
}

func buildRequest(location string) string {
	u := url.URL{}

	q := u.Query()
	q.Add("key", Utils.GetEnvVariable("WEATHER_TOKEN"))
	q.Add("q", location)
	q.Add("lang", "tr")

	u.Scheme = Utils.GetEnvVariable("WEATHER_SCHEMA")
	u.Host = Utils.GetEnvVariable("WEATHER_HOST")
	u.Path = Utils.GetEnvVariable("WEATHER_PATH")
	u.RawQuery = q.Encode()

	return u.String()
}
