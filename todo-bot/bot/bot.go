package bot

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo-bot/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			switch update.Message.Command() {
			case "add":
				_, taskText, _ := strings.Cut(update.Message.Text, update.Message.Command()+" ")

				if len(taskText) == 0 {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Task is empty.")
					bot.Send(msg)
					continue
				}

				err := database.SetTask(update.Message.From.ID, taskText)
				message := ""

				if err != nil {
					message = "Error"
				} else {
					message = "Ok"
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(msg)
				continue
			case "delete":
				_, taskText, _ := strings.Cut(update.Message.Text, update.Message.Command()+" ")

				if len(taskText) == 0 {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Task is empty.")
					bot.Send(msg)
					continue
				}

				err := database.DeleteTask(update.Message.From.ID, taskText)

				message := ""

				if err != nil {
					message = "Error"
				} else {
					message = "Ok"
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(msg)
				continue
			case "list":
				tasks, err := database.GetTaskList(update.Message.From.ID)

				message := ""

				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error")
					bot.Send(msg)
					continue
				}

				if len(tasks) > 0 {

					message = "Your tasks:\n"

					for i := 0; i < len(tasks); i++ {
						message += strconv.Itoa(i+1) + ") " + tasks[i] + "\n"
					}
				} else {
					message = "You don't have tasks"
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(msg)
				continue
			case "help":
				message := "/add - add task.\n/delete - delete task.\n/list - show all your tasks.\n/help - help."
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(msg)
				continue
			default:
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command."))
			}
		}
	}
}
