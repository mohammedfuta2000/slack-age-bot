package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents( analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
func main()  {
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A06H2K4V2UC-6590889596801-6b5fa16ceb15fc00617ddb3c9a4ba2d7e5f5ee37e0d38756417eb4e268592f4c")
	os.Setenv("SLACK_BOT_TOKEN","xoxb-6425834389143-6580753070068-JZcIPupjSHinFsbQAwsNOcs7")

	bot:=slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples: []string{"my yob is 2020","my year of birth is 2021"},
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			year:=r.Param("year")
			yob,err:=strconv.Atoi(year)
			if err!=nil{
				println(err)
			}
			age:=2023-yob
			res := fmt.Sprintf("age is %d",age)
			w.Reply(res)
		},
	})
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err:=bot.Listen(ctx)
	if err!=nil{
		log.Fatal(err)
	}
}