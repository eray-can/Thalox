package Currency

import (
	"Thalox/Pkg/Client"
	customError "Thalox/Pkg/Error"
	Currency2 "Thalox/Pkg/Model/Currency"
	"Thalox/Pkg/Utils"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var headers = map[string]string{
	"Host":             "api.genelpara.com",
	"sec-ch-ua":        "\"Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"115\", \"Chromium\";v=\"115\"",
	"sec-ch-ua-mobile": "?0",
	"Accept-Language":  "tr-TR,tr;q=0.9,en-US;q=0.8,en;q=0.7",
}

func Currency(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	url := Utils.GetEnvVariable("CURRENCY_URL")
	fmt.Println(url)
	request, err := Client.CreateRequest(url, "GET", nil, headers)
	if err != nil {
		return
	}

	response, err := Client.Response(request)
	if err != nil {
		return
	}
	var currencyData Currency2.ExchangeRates
	err = json.Unmarshal(response, &currencyData)
	if err != nil {
		return
	}
	currency, err := getCurrency(currencyData, args[1])
	if err != nil {
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       strings.ToTitle(args[1]) + " Bilgisi",
		Description: "<@" + m.Author.ID + "> " + strings.ToTitle(args[1]) + " Bilgisi aşağıda yer almaktadır",
		Color:       0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Alış",
				Value:  currency.BuyingPrice,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Satış",
				Value:  currency.SellPrice,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Değişim",
				Value:  currency.Ratio,
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://gifdb.com/images/high/dollar-bills-flipping-money-money-money-1pulbz2lornzgcp9.gif",
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}

func getCurrency(rates Currency2.ExchangeRates, currency string) (Currency2.Currency, error) {
	allRates := map[string]Currency2.Currency{
		"usd": rates.USD,
		"eur": rates.EUR,
		"gbp": rates.GBP,
		"ga":  rates.GA,
		"c":   rates.C,
		"gag": rates.GAG,
		"btc": rates.BTC,
		"eth": rates.ETH,
	}
	res, ok := allRates[currency]
	if !ok {
		err := customError.NewError(0, "currency")

		return Currency2.Currency{}, err
	}
	return res, nil
}
