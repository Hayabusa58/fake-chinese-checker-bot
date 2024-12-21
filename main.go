package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	tk := os.Getenv("BOT_TOKEN")

	session, err := discordgo.New("Bot " + tk)

	if err != nil {
		fmt.Println(err)
	}

	// onLogin event handler
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Logged in as %s", r.User.String())
	})

	// onMessageCreate event handler
	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		cid := os.Getenv("CHANNEL_ID")

		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.ChannelID == cid {
			// fmt.Println("Message received at ", cid, ": ", m.Content)
			if !IsAllowedCharacter(m.Content) {
				// fmt.Println("検知対象: ", m.Content)
				s.ChannelMessageSend(m.ChannelID, "🤖 不正入国者検知!")
			} else {
				// fmt.Println("検知対象外: ", m.Content)
				return
			}
		}

	})

	err = session.Open()

	if err != nil {
		fmt.Println(err)
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	<-sigch

	err = session.Close()
	if err != nil {
		fmt.Println(err)
	}

}
