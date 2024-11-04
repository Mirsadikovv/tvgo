package dto

import (
	"github.com/Mirsadikovv/tvgo/playlist/model"
	"github.com/fobus1289/ufa_shared/http/response"
)

type PagePlaylistResponseType = response.PaginateResponse[*model.PlaylistModel] // @name PagePlaylistResponseType

type CreatePlaylistDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LogoLink    string `json:"logoLink"`
	ChannelLink string `json:"channelLink"`
} //@name CreatePlaylistDto

type UpdatePlaylistDto struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	LogoLink    *string `json:"logoLink"`
	ChannelLink *string `json:"channelLink"`
	IsVisible   *bool   `json:"isVisible"`
} //@name UpdatePlaylistDto
