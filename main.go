package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {

	// .env load
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	kBotToken := os.Getenv("BOT_TOKEN")
	kGroupID, _ := strconv.ParseInt(os.Getenv("GROUP_ID"), 10, 64)

	// bot init
	bot, err := tgbotapi.NewBotAPI(kBotToken)
	if err != nil {
		log.Fatal("Err: ", err)
	}

	// bot.Debug = true

	// _, err = bot.SetWebhook(tgbotapi.NewWebhook("https://1b22-110-224-131-152.ngrok.io"))
	// if err != nil {
	// 	log.Fatal("Err: ", err)
	// }
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal("Err: ", err)
	}
	fmt.Println(info.URL)
	updates := bot.ListenForWebhook("/")

	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	for update := range updates {
		log.Printf("%+v\n", update.Message.Text)

		if update.Message.IsCommand() {
			isDefault := false
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "start":
				msg.Text = "Use this bot to echo some shit.\n/start - Info\n/echo - Echo stuff normally\n/echob - Echo stuff in bold\n/echoi - Echo stuff in italic\n/echou - Echo stuff underlined\n/echom - Echo stuff in monospace\n/echo(b/i/u) - Echo stuff in combinations of b/i/u, order alphabetically"
			case "echo":
				msg.Text = update.Message.CommandArguments()
			case "echob":
				msg.ParseMode = "MarkdownV2"
				msg.Text = "*" + update.Message.CommandArguments() + "*"
			case "echoi":
				msg.ParseMode = "MarkdownV2"
				msg.Text = "_" + update.Message.CommandArguments() + "_"
			case "echou":
				msg.ParseMode = "MarkdownV2"
				msg.Text = "__" + update.Message.CommandArguments() + "__"
			case "echobi":
				msg.ParseMode = "MarkdownV2"
				msg.Text = "*_" + update.Message.CommandArguments() + "_*"
			case "echobu":
				msg.ParseMode = "MarkdownV2"
				msg.Text = "*__" + update.Message.CommandArguments() + "__*"
			case "echoiu":
				msg.ParseMode = "MarkdownV2"
				msg.Text = "_ __" + update.Message.CommandArguments() + "__ _"
			case "echobiu":
				msg.ParseMode = "MarkdownV2"
				msg.Text = "*_ __" + update.Message.CommandArguments() + "__ _*"
			case "echom":
				msg.ParseMode = "MarkdownV2"
				msg.Text = "`" + update.Message.CommandArguments() + "`"
			default:
				isDefault = true
			}
			if isDefault {
				continue
			}
			if update.Message.Chat.ID != kGroupID {
				msg.Text = "`[!] Bot Not Authorised`"
				msg.ParseMode = "MarkdownV2"
			}
			bot.Send(msg)
		}
	}
}
