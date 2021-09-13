package repository

import (
	"net"

	"github.com/wellWINeo/GoIpBot"
	"gorm.io/gorm"
)

// interface for interaction with
// IpStack API
type IpStack interface {
	GetInfo(net.IP) (GoIpBot.IpInfo, error)
}

type User interface {
	AddHistory(int64, net.IP, GoIpBot.IpInfo) error
	GetHistory(int64) ([]GoIpBot.HistoryRecord, error)
}

type Admin interface {
	GrantAdmin(string) error
	RevokeAdmin(string) error
	GetUsers() ([]GoIpBot.User, error)
	GetUserHistory(string) ([]GoIpBot.HistoryRecord, error)
	CheckPermission(int64) bool
	GetUser(int) (GoIpBot.User, error)
	RemoveHistory(int) error
}

type Repository struct {
	IpStack
	User
	Admin
}

func NewRepository(token string, db *gorm.DB) *Repository {
	return &Repository{
		IpStack: NewIpStack(token),
		User:    NewUserGORM(db),
		Admin:   NewAdminGORM(db),
	}
}
