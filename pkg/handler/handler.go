package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/wellWINeo/GoIpBot"
	"github.com/wellWINeo/GoIpBot/pkg/service"
)

// Backend
type WebHandler struct {
	services *service.Service
}

func NewWebHandler(services *service.Service) *WebHandler {
	return &WebHandler{services: services}
}

func (t *WebHandler) InitRoutes(port int) error {
	http.HandleFunc("/get_users", t.GetUsers)
	http.HandleFunc("/get_user", t.GetUser)
	http.HandleFunc("/get_history_by_tg", t.GetHistoryByTg)
	http.HandleFunc("/remove_history", t.RemoveHistory)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

// Telegram bot functions
type TeleHandler struct {
	services *service.Service
	handlers map[string]GoIpBot.TelegramMessageHandler
}

func NewTeleHandler(services *service.Service) *TeleHandler {
	return &TeleHandler{
		services: services,
		handlers: map[string]GoIpBot.TelegramMessageHandler{},
	}
}

func (t *TeleHandler) InitRoutes(in chan tgbotapi.Message, out chan tgbotapi.MessageConfig) {

	// fill routes table
	t.handlers["/ip"] = t.GetInfo
	t.handlers["/history"] = t.GetHistory
	t.handlers["/grant"] = t.GranAdmin
	t.handlers["/revoke"] = t.RevokeAdmin
	t.handlers["/users"] = t.GetUsers
	t.handlers["/user_history"] = t.GetUserHistory
	t.handlers["/help"] = t.Help

	// returns routing function
	go func () {
		for msg := range in {
			var (
				str string
				err error
				resp tgbotapi.MessageConfig
			)

			GoIpBot.Log("handler.go").Info("message received to process")

			splittedMsg := strings.Split(msg.Text, " ")
			if  t.handlers[splittedMsg[0]] != nil {
				str, err = t.handlers[splittedMsg[0]](msg.Chat.ID, msg.From.UserName,
					splittedMsg[1:])
			} else if splittedMsg[0] == "/broadcast"{
				var msgs []tgbotapi.MessageConfig
				msgs, err = t.Broadcast(msg.Chat.ID, msg.Chat.UserName,
					splittedMsg[1:])
				if err == nil {
					GoIpBot.Log("handler.go").Info("broadcasting message")
					for _, m := range msgs {
						out <- m
					}
					continue
				}
			} else {
				err = errors.New("To many words")
				GoIpBot.Log("handler.go").Error(err)
			}

			if err != nil {
				str = fmt.Sprintf("[Error] %s", err.Error())
			}
			resp = tgbotapi.NewMessage(msg.Chat.ID, str)

			GoIpBot.Log("handler.go").Info("sending message to out queue")

			out <- resp
		}
	}()
}
