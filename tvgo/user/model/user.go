package model

import (
	"time"
)

type UserModel struct {
	Id               int64      `json:"id" gorm:"primaryKey"`
	FullName         string     `json:"fullName" gorm:"unique"`
	Username         string     `json:"username" gorm:"unique"`
	TelegramId       int        `json:"telegramId" gorm:"unique"`
	RegTime          *time.Time `json:"regTime"`
	Phone            string     `json:"phone" gorm:"unique"`
	Status           string     `json:"status"`
	SubscriptionDate *time.Time `json:"subscriptionDate"`
	IsVisible        bool       `json:"isVisible" gorm:"default:true"`
	CreatedAt        *time.Time `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt        *time.Time `json:"updatedAt,omitempty" gorm:"autoUpdateTime:true"`
}

func (UserModel) TableName() string {
	return "users"
}
