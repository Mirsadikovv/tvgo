package dto

import (
	"time"

	"github.com/Mirsadikovv/tvgo/user/model"
	"github.com/fobus1289/ufa_shared/http/response"
)

type PageUserResponseType = response.PaginateResponse[*model.UserModel] // @name PageUserResponseType

type CreateUserDto struct {
	FullName         string     `json:"fullName"`
	Username         string     `json:"username"`
	TelegramId       int        `json:"telegramId"`
	RegTime          *time.Time `json:"regTime"`
	Phone            string     `json:"phone"`
	Status           string     `json:"status"`
	SubscriptionDate *time.Time `json:"subscriptionDate"`
} //@name CreateUserDto

type UpdateUserDto struct {
	FullName         *string    `json:"fullName"`
	Username         *string    `json:"username"`
	TelegramId       *int       `json:"telegramId"`
	Phone            *string    `json:"phone"`
	Status           *string    `json:"status"`
	SubscriptionDate *time.Time `json:"subscriptionDate"`
	IsVisible        *bool      `json:"isVisible"`
} //@name UpdateUserDto
