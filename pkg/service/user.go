package service

import (
	"net"

	"github.com/wellWINeo/GoIpBot"
	"github.com/wellWINeo/GoIpBot/pkg/repository"
)

type UserService struct {
	repo repository.User
	api  repository.IpStack
}

func NewUser(repo repository.User, api repository.IpStack) *UserService {
	return &UserService{
		repo: repo,
		api: api,
	}
}

func (u *UserService) GetInfo(userID int64, ip net.IP) (GoIpBot.IpInfo, error) {
	info, err := u.api.GetInfo(ip)
	if err != nil {
		return info, err
	}

	err = u.repo.AddHistory(userID, ip, info)

	return info, err
}

func (u *UserService) GetHistory(userID int64) ([]GoIpBot.HistoryRecord, error) {
	history, err := u.repo.GetHistory(userID)
	if err != nil {
		return nil, err
	}

	// slice to return
	response := make([]GoIpBot.HistoryRecord, len(history))
	seen := make(map[string]bool)

	for i := len(history) - 1; i >= 0; i-- {
		if !seen[history[i].Address] {
			response = append(response, history[i])
			seen[history[i].Address] = true
		}
	}

	return response, nil
}
