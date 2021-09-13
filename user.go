package GoIpBot

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TelegramID int64           `gorm:"telegram_id" json:"telegram_id"`
	TagName    string          `gorm:"tag_name" json:"tag_name"`
	IsAdmin    bool            `gorm:"is_admin" json:"is_admin"`
	History    []HistoryRecord `gorm:"foreignKey:UserNumber;"`
}
