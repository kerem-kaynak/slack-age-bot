package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6395536233797-6398381389235-YhDIQ3gK3olDgnHBoW7BuKV2")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06B9TFGQG7-6421255108960-97d43a12d73a544004af824efd9a53ebf570f520ec6f938715761a0b4044da2b")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My yob is year <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"My yob is year 1998"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			fmt.Printf("year: %s\n", year)
			yob, err := strconv.Atoi(year)
			fmt.Printf("yob: %d\n", yob)
			if err != nil {
				fmt.Println("error")
			}
			age := 2024 - yob
			fmt.Printf("age: %d\n", age)
			r := fmt.Sprintf("Your age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
