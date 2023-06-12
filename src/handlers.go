package src

import (
	"fmt"
	"log"
	"strings"

	"github.com/amarnathcjd/gogram/telegram"
)

func PMStartHandler(message *telegram.NewMessage) error {
	message.Reply("Hello, I am an Account Manager Bot!")
	return nil
}

func NewAccountHandler(message *telegram.NewMessage) error {
	conv, err := message.Client.NewConversation(message.Sender, false, 600)
	if err != nil {
		return err
	}
	defer conv.Close()
	conv.SendMessage("Enter the phone number of the account you want to add")
	phone, err := conv.GetResponse()
	if err != nil {
		return err
	}
	conv.SendMessage("Enter the app id of the account you want to add")
	appid, err := conv.GetResponse()
	if err != nil {
		return err
	}
	conv.SendMessage("Enter the app hash of the account you want to add")
	apphash, err := conv.GetResponse()
	if err != nil {
		return err
	}
	log.Println("Logging in to the account", phone.Text(), appid.Text(), apphash.Text())
	newclient, err := telegram.NewClient(telegram.ClientConfig{
		AppID:         Converttoin32(appid.Text()),
		AppHash:       apphash.Text(),
		LogLevel:      telegram.LogInfo,
		MemorySession: true,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := newclient.Connect(); err != nil {
		return err
	}
	code, err := newclient.SendCode(phone.Text())
	if err != nil {
		return err
	}
	conv.SendMessage("Enter the OTP of the account you want to add in the format 1 2 3 4 5")
	otp, err := conv.GetResponse()
	if err != nil {
		return err
	}
	if _, err := newclient.AuthSignIn(strings.ReplaceAll(phone.Text(), " ", ""), code, strings.ReplaceAll(otp.Text(), " ", ""), nil); err != nil {
		if strings.Contains(err.Error(), "SESSION_PASSWORD_NEEDED") {
			conv.SendMessage("Enter the 2FA password of the account you want to add")
			password, _ := conv.GetResponse()
			pass, _ := newclient.AccountGetPassword()
			inputpass, err := telegram.GetInputCheckPassword(password.Text(), pass)
			if err != nil {
				return err
			}
			_, err = newclient.AuthCheckPassword(inputpass)
			if err != nil {
				message.Reply("Invalid Password")
				return err
			}
		} else {
			message.Reply(fmt.Sprintf("Error logging in to the account %s", err.Error()))
			return err
		}
	}

	user, _ := newclient.GetMe()
	go AddAccount(user.ID, message.SenderID(), Converttoin32(appid.Text()), apphash.Text(), newclient.ExportSession())
	message.Reply(fmt.Sprintf("Account <b>%s</b> added successfully", user.FirstName))
	return nil
}

func GetUserAccountsHandler(message *telegram.NewMessage) error {
	accounts, err := FetchAccounts(message.SenderID())
	if err != nil {
		return err
	}
	if len(accounts) == 0 {
		message.Reply("You have no accounts added")
		return nil
	}
	var reply string
	for i, account := range accounts {
		fmt.Println(account.UserID, i+1)
		reply += fmt.Sprintf("<b>➤</b> <a href='tg://user?id=%d'>Account %d</a>\n", account.UserID, i+1)
	}
	_, err = message.Reply(reply)
	if err != nil {
		return err
	}
	return nil
}

func SendMsg(message *telegram.NewMessage) error {
	accounts, err := FetchAccounts(message.SenderID())
	if err != nil {
		return err
	}
	if len(accounts) == 0 {
		message.Reply("You have no accounts added")
		return nil
	
	}
	message.Reply(accounts[0].StringSession)
	fmt.Println(accounts[0].StringSession)
	fmt.Println(len(strings.TrimSpace(accounts[0].StringSession)))
	clients, err := LoginClient(accounts)
	if err != nil {
		message.Reply("Error logging in to the accounts")
		fmt.Println(err)
		return err
	}
	
	for _, client := range clients {
		client.SendMessage("tamilvip07", "Hello")
	}
	return nil
}
