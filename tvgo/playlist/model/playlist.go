package model

import (
	"time"

	"gorm.io/gorm"
)

type PlaylistModel struct {
	Id          int64           `json:"id" gorm:"primaryKey"`
	Name        string          `json:"name" gorm:"unique"`
	Description string          `json:"description"`
	LogoLink    string          `json:"logoLink"`
	ChannelLink string          `json:"channelLink"`
	IsVisible   bool            `json:"isVisible" gorm:"default:true"`
	CreatedAt   *time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt   *time.Time      `json:"updatedAt,omitempty" gorm:"autoUpdateTime:true"`
	DeletedAt   *gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

func (PlaylistModel) TableName() string {
	return "playlists"
}
