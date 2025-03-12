package dto

import (
	"github.com/Mirsadikovv/tvgo/playlist/model"
	"github.com/fobus1289/ufa_shared/http/response"
)

type ChannelType string // @name ChannelType

const (
	VIP   ChannelType = "VIP"
	SPORT ChannelType = "SPORT"
	FREE  ChannelType = "FREE"
)

type PagePlaylistResponseType = response.PaginateResponse[*model.PlaylistModel] // @name PagePlaylistResponseType

type CreatePlaylistDto struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	LogoLink    string      `json:"logoLink"`
	ChannelLink string      `json:"channelLink"`
	Type        ChannelType `json:"type"`
} //@name CreatePlaylistDto

type UpdatePlaylistDto struct {
	Name        *string      `json:"name"`
	Description *string      `json:"description"`
	LogoLink    *string      `json:"logoLink"`
	ChannelLink *string      `json:"channelLink"`
	IsVisible   *bool        `json:"isVisible"`
	Type        *ChannelType `json:"type"`
} //@name UpdatePlaylistDto

type PlaylistQueryParams struct {
	Name      *string      `query:"name"`
	IsVisible *bool        `query:"isVisible"`
	Type      *ChannelType `query:"type"`
} //@name PlaylistQueryParams
