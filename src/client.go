package src

import (
	"fmt"

	"github.com/amarnathcjd/gogram/telegram"
)

var bot *telegram.Client

func InitBot() *telegram.Client {
	// Create a new client
	bot, _ := telegram.NewClient(telegram.ClientConfig{
		AppID:    Envars.AppId,
		AppHash:  Envars.AppHash,
		LogLevel: telegram.LogInfo,
		Session:  "./bot.session",
	})
	if err := bot.Connect(); err != nil {
		panic(err)
	}
	if err := bot.LoginBot(Envars.Token); err != nil {
		panic(err)
	}
	bot.AddMessageHandler("/start", PMStartHandler)
	bot.AddMessageHandler("/addacc", NewAccountHandler)
	bot.AddMessageHandler("/listacc", GetUserAccountsHandler)
	bot.AddMessageHandler("/send", SendMsg)
	return bot
}

func LoginClient(accs []Accounts) ([]*telegram.Client, error) {
	clients := make([]*telegram.Client, 0)
	for _, account := range accs {
		client, err := telegram.NewClient(telegram.ClientConfig{
			AppID:         account.AppId,
			AppHash:       account.AppHash,
			LogLevel:      telegram.LogInfo,
			StringSession: account.StringSession,
			MemorySession: true,
		})
		if err != nil {
			fmt.Println("Error creating new client: ", err)
			return nil, err
		}
		if err := client.Start(); err != nil {
			fmt.Println("Error starting client: ", err)
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}
