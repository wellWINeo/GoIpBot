package repository

import (
	"errors"

	"github.com/wellWINeo/GoIpBot"
	"gorm.io/gorm"
)

type AdminGORM struct {
	db *gorm.DB
}

func NewAdminGORM(db *gorm.DB) *AdminGORM {
	return &AdminGORM{db: db}
}

// set user as admin
func (a *AdminGORM)	GrantAdmin(tag string) error {
	// check that user with such tag in db
	if a.db.Model(&GoIpBot.User{}).Where("tag_name = ?", tag).
		First(&GoIpBot.User{}).Error != nil {
		GoIpBot.Log("admin.go").Error("No such user")
		return errors.New("No such user")
	}

	result := a.db.Model(&GoIpBot.User{}).Where("tag_name = ?", tag).
		Update("is_admin", true)
	if result.Error != nil {
		GoIpBot.Log("admin.go").Error(result.Error)
	}
	return result.Error
}

// remove admin priviliges
func (a *AdminGORM)	RevokeAdmin(tag string) error {
	// check that user with such tag in db
	if a.db.Model(&GoIpBot.User{}).Where("tag_name = ?", tag).
		First(&GoIpBot.User{}).Error != nil {
		GoIpBot.Log("admin.go").Error("No such user")
		return errors.New("No such user")
	}

	result := a.db.Model(&GoIpBot.User{}).Where("tag_name = ?", tag).
		Update("is_admin", false)
	if result.Error != nil {
		GoIpBot.Log("admin.go").Error(result.Error)
	}
	return result.Error
}

// get all users
func (a *AdminGORM)	GetUsers() ([]GoIpBot.User, error) {
	var users []GoIpBot.User
	result := a.db.Model(&GoIpBot.User{}).Find(&users)
	if result.Error != nil {
		GoIpBot.Log("admin.go").Error(result.Error)
	}
	return users, result.Error
}

// get user's history
func (a *AdminGORM)	GetUserHistory(tag string) ([]GoIpBot.HistoryRecord, error) {
	var user GoIpBot.User
	result := a.db.Model(&user).Where("tag_name = ?", tag).Find(&user)
	if result.Error != nil {
		GoIpBot.Log("admin.go").Error(result.Error)
		return []GoIpBot.HistoryRecord{}, result.Error
	} else if user.ID == 0 {
		GoIpBot.Log("admin.go").Error("No such uer")
		return []GoIpBot.HistoryRecord{}, errors.New("No such user")
	}

	var history []GoIpBot.HistoryRecord
	result = a.db.Model(&GoIpBot.HistoryRecord{}).Where("user_number = ?",
		user.ID).Find(&history)
	if result.Error != nil {
		GoIpBot.Log("admin.go").Error(result.Error)
	}

	return history, result.Error
}

// check that user is admin
func (a *AdminGORM) CheckPermission(chatID int64) bool {
	var user GoIpBot.User
	result := a.db.Where("telegram_id = ?", chatID).First(&user)
	if result.Error != nil {
		GoIpBot.Log("admin.go").Error(result.Error)
		return false
	}

	return user.IsAdmin
}

// get user by ID in db
func (a *AdminGORM) GetUser(id int) (GoIpBot.User, error) {
	var user GoIpBot.User
	result := a.db.First(&user, "id = ?", id)
	if result.Error != nil {
		GoIpBot.Log("admin.go").Error(result.Error)
	}
	return user, result.Error
}

// remove history record from db
func (a *AdminGORM) RemoveHistory(id int) error {
	result := a.db.Delete(&GoIpBot.HistoryRecord{}, id)
	if result.Error != nil {
		GoIpBot.Log("admin.go").Error(result.Error)
	}
	return result.Error
}
