package handler

import (
	"fmt"
	"net"
)

// handlers must confirm TelegramMessageHandler type

func (t *TeleHandler) Help(userID int64, tag string, args []string) (string, error) {
	response := fmt.Sprint("Bot to check IP address\n" +
		"/ip <addres> - checking IP address\n" +
		"/history - return history of your queries\n" +
		"/help - view this message\n" +
		"Admin functions:\n" +
		"/grant <username> - add user to admins\n" +
		"/revoke <username> - remove user from admins\n" +
		"/users - list all users\n" +
		"/user_history <username> - get user's history")
	return response, nil
}

func (t *TeleHandler) GetInfo(userID int64, tag string, args []string) (string, error) {
	if len(args) != 1 {
		return "", t.errorResponse("Wrong arguments count")
	}

	ip := net.ParseIP(args[0])
	if ip == nil {
		return "", t.errorResponse("Can't parse IP address")
	}

	info, err := t.services.GetInfo(userID, ip)
	if err != nil {
		return "", t.errorResponse(err.Error())
	}

	return fmt.Sprintf("Country: %s", info.CountryName), nil
}

func (t *TeleHandler) GetHistory(userID int64, tag string, args []string) (string, error) {
	var response string
	history, err := t.services.GetHistory(userID)
	if err != nil {
		return "", t.errorResponse(err.Error())
	}

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
