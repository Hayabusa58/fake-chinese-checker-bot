package main

import (
	"fmt"
	"os"
	"os/signal"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type config struct {
	Token     string `env:"BOT_TOKEN"`
	ChannelID string `env:"CHANNEL_ID"`
}

// ToDo: emoji と漢字が混合した文字列は除外できない
func containsTargetCharacters(s string) bool {
	// emoji 単体は除外する
	emojireg := regexp.MustCompile(`^<:[^:>]+:[0-9]+>$`)
	re := regexp.MustCompile(`[a-zA-Z0-9ァ-ヶぁ-ん]`)
	return re.MatchString(s) && !emojireg.MatchString(s)
}

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
			fmt.Println("Message received at ", cid, ": ", m.Content)
			if containsTargetCharacters(m.Content) {
				s.ChannelMessageSend(m.ChannelID, "🤖 不正入国者検知!")
			} else {
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
