package Handle

import (
	"Thalox/Pkg/Command/Currency"
	"Thalox/Pkg/Command/Help"
	"Thalox/Pkg/Command/Weather"
	"Thalox/Pkg/Command/Who"
	customError "Thalox/Pkg/Error"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var commands = map[int32]map[string]func(botSession *discordgo.Session, m *discordgo.MessageCreate, args []string){
	0: {
		"kimsin": Who.Who,
		"yardım": Help.Help,
	},
	1: {
		"kur":  Currency.Currency,
		"hava": Weather.Weather,
	},
}

func MainHandle(botSession *discordgo.Session, m *discordgo.MessageCreate) {
	message, status, err := messageHandle(m.Content, m.Author.ID, botSession.State.User.ID, m.Author.Bot)
	if err != nil {
		return
	}

	if cmdFunc, found := commands[status][message[0]]; found {
		cmdFunc(botSession, m, message)
	}
}

func messageHandle(userMessage string, AuthId string, botId string, isBot bool) ([]string, int32, error) {
	err := customError.NewError(0, userMessage)

	if AuthId == botId || isBot {
		return nil, 0, customError.UpdateError(1, err).DeveloperMessage
	}

	if !strings.HasPrefix(userMessage, ">") {
		return nil, 0, customError.UpdateError(2, err).DeveloperMessage
	}

	replaceMessage := strings.Replace(strings.ToLower(userMessage), ">", "", 1)
	message := strings.Split(replaceMessage, " ")

	return message, int32(len(message) - 1), nil
}
