package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"astrobot/sdk/client"
	"astrobot/sdk/client/operations"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func main() {
	var Token string
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

	// Create a new Discordgo session
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println(err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	err = dg.Open()
	if err != nil {
		panic("error while opening connection")
	}
	defer dg.Close()

	go func() {
		for {
			user := "NicolasF"
			response, err := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{Host: "localhost:3000"}).Operations.GetImages(&operations.GetImagesParams{
				User:       &user,
				Context:    context.Background(),
				HTTPClient: &http.Client{Timeout: 5 * time.Second},
			})
			if err != nil {
				logrus.Error(err)
				continue
			}

			dg.ChannelMessageSend("1053456333145378858", *response.Payload[0].Title)
			time.Sleep(5 * time.Second)
		}
	}()

	fmt.Println("Bot is now running, press CTRL C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	s.ChannelMessageSend(m.ChannelID, "test")
}
