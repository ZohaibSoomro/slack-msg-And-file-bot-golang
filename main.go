package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
	"github.com/zohaibsoomro/myslackbot/config"
)

const (
	Bot_Token  = "Bot_Token"
	App_Token  = "App_Token"
	Channel_Id = "Channel_Id"
)

func main() {
	config.SetEnvironment()
	client := slacker.NewClient(os.Getenv(Bot_Token), os.Getenv(App_Token))

	cmdDef := &slacker.CommandDefinition{
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			w.Reply("Hey! I am your slack bot assistant, how may i help you?")
		},
	}
	client.Command("hi", cmdDef)
	client.Command("Hello", cmdDef)
	client.Command("where is file?", &slacker.CommandDefinition{
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			params := slack.FileUploadParameters{
				File:     "./file.txt",
				Title:    "file",
				Channels: []string{os.Getenv(Channel_Id)},
			}
			cl := slack.New(os.Getenv(Bot_Token), slack.OptionAppLevelToken(os.Getenv(App_Token)))
			w.Reply("Here is your file")
			file, err := cl.UploadFile(params)
			if err != nil {
				w.Reply("Some error while sending file.")
			}
			fmt.Println(file.Name)
		},
	})
	go printCommandEvents(client.CommandEvents())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Fatal(client.Listen(ctx))
	fmt.Scan(new(int))
}

func printCommandEvents(ch <-chan *slacker.CommandEvent) {
	for event := range ch {
		fmt.Println(event.Command)
		fmt.Println(event.Timestamp)
		fmt.Println(event.Event.Text)
		fmt.Println(event.Parameters)
	}
}
