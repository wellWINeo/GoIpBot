package handler

import (
	"errors"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/wellWINeo/GoIpBot"
)

func (t *TeleHandler) errorResponse(msg string) error {
	err := errors.New(msg)
	GoIpBot.Log("tele_*.go").Error(err)
	return err
}

// handlers must confirm TelegramMessageHandler type

func (t *TeleHandler) GranAdmin(userID int64, tag string, args []string) (string, error) {
	if !t.services.CheckPermission(userID) {
		return "", t.errorResponse("permission denied")
	}

	if len(args) != 1 {
		return "", t.errorResponse("wrong argument count")
	}

	err := t.services.GrantAdmin(args[0])

	if err != nil {
		return "", t.errorResponse(err.Error())
	} else {
		return "User added to admins", nil
	}
}

func (t *TeleHandler) RevokeAdmin(userID int64, tag string, args []string) (string, error) {
	if !t.services.CheckPermission(userID) {
		return "", t.errorResponse("permission denied")
	}

	if len(args) != 1 {
		return "", t.errorResponse("Wrong argument count")
	}

	err := t.services.RevokeAdmin(args[0])

	if err != nil {
		return "", t.errorResponse(err.Error())
	} else {
		return "User removed from admins", nil
	}
}

func (t *TeleHandler) GetUsers(userID int64, tag string, args []string) (string, error) {
	if !t.services.CheckPermission(userID) {
		return "", t.errorResponse("permission denied")
	}

	var resp string

	users, err := t.services.Admin.GetUsers()
	if err != nil {
		return "", t.errorResponse(err.Error())
	}

	for _, user := range users {
		resp += fmt.Sprintf("ID: %d\nTag: %s\nChat ID: %d\n Admin: %t",
			user.ID, user.TagName, user.TelegramID, user.IsAdmin)
		resp += "---"
	}

	return resp, nil
}

func (t *TeleHandler) GetUserHistory(userID int64, tag string, args []string) (string, error) {
	if !t.services.CheckPermission(userID) {
		return "", t.errorResponse("permission denied")
	}

	if len(args) != 1 {
		return "", t.errorResponse("Wrong argument count")
	}

	history, err := t.services.GetUserHistory(args[0])
	if err != nil {
		return "", t.errorResponse(err.Error())
	}

	var response string

	if len(history) > 0 {
		for _, h := range history {
			response += fmt.Sprintf("[%v] - %s - %s\n", h.CreatedAt,
				h.Address, h.CountryName)
		}
	} else {
		response = "Empty ;)"
	}

	return response, nil
}

// Broadcast function
func (t *TeleHandler) Broadcast(userID int64, tag string, args[] string) ([]tgbotapi.MessageConfig, error) {
	if !t.services.CheckPermission(userID) {
		return nil, t.errorResponse("permission denied")
	}

	if len(args) == 0 {
		return nil, errors.New("No argument to send")
	}

	text := fmt.Sprintf("From: %s\n", tag)
	text += strings.Join(args, " ")

	users, err := t.services.GetUsers()
	if err != nil {
		return nil, err
	}

	messages := make([]tgbotapi.MessageConfig, len(users))

	for _, user := range users {
		messages = append(messages, tgbotapi.NewMessage(user.TelegramID, text))
	}

	return messages, nil
}
