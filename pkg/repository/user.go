package repository

import (
	"net"

	"github.com/wellWINeo/GoIpBot"
	"gorm.io/gorm"
)

type UserGORM struct {
	db *gorm.DB
}

func NewUserGORM(db *gorm.DB) *UserGORM {
	return &UserGORM{db: db}
}

func (u *UserGORM) GetHistory(userID int64) ([]GoIpBot.HistoryRecord, error) {
	var user GoIpBot.User
	result := u.db.Model(&user).Where("telegram_id = ?", userID).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	var history []GoIpBot.HistoryRecord
	result = u.db.Model(&GoIpBot.HistoryRecord{}).Where("user_number = ?",
		user.ID).Find(&history)

	return history, result.Error
}

func (u *UserGORM) CreateUser(userID int64) (int, error) {
	user := GoIpBot.User{
		Model:      gorm.Model{},
		TelegramID: userID,
		IsAdmin:    false,
		History:    []GoIpBot.HistoryRecord{},
	}
	result := u.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	result = u.db.Where("telegram_id = ?", userID).Find(&user)
	return int(user.ID), result.Error
}

func (u *UserGORM) AddHistory(userID int64, ip net.IP, info GoIpBot.IpInfo) error {
	history := GoIpBot.HistoryRecord{
		Model:      gorm.Model{},
		Address:    ip.String(),
		IpInfo:     info,
		UserNumber: userID,
	}

	var user GoIpBot.User
	result := u.db.First(&user, "telegram_id = ?", userID)
	if result.Error == gorm.ErrRecordNotFound {
		id, err := u.CreateUser(userID)
		if err != nil {
			return err
		}
		result := u.db.First(&user, id)

		if result.Error != nil {
			return result.Error
		}
	} else if result.Error != nil {
		return result.Error
	}

	user.History = append(user.History, history)

	result = u.db.Save(&user)
	return result.Error
}
