package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent)  {
	for event :=range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()


	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN","xoxb-5430346750560-5403750329781-eRSfiPWf64phoYi3yAfyPiaJ")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A05CNASC1LG-5406663275362-65594f425c6ed353939d0ab0a0c822ce04adf0986e2d63f5938e6127797f0105")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		// Example: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year :=request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2023-yob
			r := fmt.Sprintf("age is %d",age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err :=bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}

}