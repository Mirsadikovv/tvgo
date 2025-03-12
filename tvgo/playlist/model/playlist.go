package model

import (
	"time"
)

type PlaylistModel struct {
	Id          int64      `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name" gorm:"unique"`
	Description string     `json:"description"`
	LogoLink    string     `json:"logoLink"`
	ChannelLink string     `json:"channelLink"`
	Type        string     `json:"type"`
	IsVisible   bool       `json:"isVisible" gorm:"default:true"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty" gorm:"autoUpdateTime:true"`
}

func (PlaylistModel) TableName() string {
	return "playlists"
}
