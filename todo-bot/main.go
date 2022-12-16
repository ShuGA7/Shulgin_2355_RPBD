package main

import (
	"fmt"
	"todo-bot/bot"
	"todo-bot/database"
)

func main() {
	fmt.Println("start")
	err := database.CheckTable()

	if err == nil {
		bot.Start()
	}

	fmt.Println("stop")
}
