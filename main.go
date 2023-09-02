package main

import (
	"Thalox/Pkg/Handle"
	"Thalox/Pkg/Utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
)

func main() {
	if err := godotenv.Load(); err != nil {
		return
	}
	token := Utils.GetEnvVariable("BOT_TOKEN")

	fmt.Println(token)
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("bot oluşturma hatası")
		return
	}

	open := bot.Open()
	if open != nil {
		fmt.Println("Bot başlama hatası")
		return
	}

	bot.AddHandler(Handle.MainHandle)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc

}
