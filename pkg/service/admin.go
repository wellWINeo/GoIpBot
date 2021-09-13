package service

import (
	"errors"

	"github.com/wellWINeo/GoIpBot"
	"github.com/wellWINeo/GoIpBot/pkg/repository"
)

type AdminService struct {
	repo repository.Admin
}

func NewAdmin(repo repository.Admin) *AdminService {
	return &AdminService{repo: repo}
}

func (a *AdminService) GrantAdmin(tag string) error {
	return a.repo.GrantAdmin(tag)
}

func (a *AdminService) RevokeAdmin(tag string) error {
	return a.repo.RevokeAdmin(tag)
}

func (a *AdminService) BroadCast(chatID int64, msg string) error {
	return errors.New("not implemented")
}

func (a *AdminService) GetUsers() ([]GoIpBot.User, error) {
	return a.repo.GetUsers()
}

func (a *AdminService) GetUserHistory(tag string) ([]GoIpBot.HistoryRecord, error) {
	return a.repo.GetUserHistory(tag)
}

func (a *AdminService) CheckPermission(chatID int64) bool {
	return a.repo.CheckPermission(chatID)
}

func (a *AdminService) GetUser(id int) (GoIpBot.User, error) {
	return a.repo.GetUser(id)
}

func (a *AdminService) RemoveHistory(id int) error {
	return a.repo.RemoveHistory(id)
}
