package service

import (
	"net"

	"github.com/wellWINeo/GoIpBot"
	"github.com/wellWINeo/GoIpBot/pkg/repository"
)

// interface for Admin's functions
type Admin interface {
	GrantAdmin(tag string) error
	RevokeAdmin(tag string) error
	BroadCast(chatID int64, msg string) error
	GetUsers() ([]GoIpBot.User, error)
	GetUserHistory(tag string) ([]GoIpBot.HistoryRecord, error)
	CheckPermission(chatID int64) bool
	GetUser(id int) (GoIpBot.User, error)
	RemoveHistory(id int) error
}

// interface for user (public)
type User interface {
	GetInfo(chatID int64, ip net.IP) (GoIpBot.IpInfo, error)
	GetHistory(chatID int64) ([]GoIpBot.HistoryRecord, error)
}

type Service struct {
	Admin
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Admin: NewAdmin(repos.Admin),
		User: NewUser(repos.User, repos.IpStack),
	}
}
