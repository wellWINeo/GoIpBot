package GoIpBot

import (
	"gorm.io/gorm"
)

type HistoryRecord struct {
	gorm.Model
	Address string `gorm:"address"`
	IpInfo
	UserNumber int64
}
