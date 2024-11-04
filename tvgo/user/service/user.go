package service

import (
	"context"
	"math"

	"github.com/Mirsadikovv/tvgo/user/dto"
	"github.com/Mirsadikovv/tvgo/user/model"
	"github.com/fobus1289/ufa_shared/http/response"
	"gorm.io/gorm"
)

type ServiceScope = func(d *gorm.DB) *gorm.DB

type UserService interface {
	FindOne(ctx context.Context, scopes ...ServiceScope) (*model.UserModel, error)
	Find(ctx context.Context, scopes ...ServiceScope) ([]model.UserModel, error)
	Page(ctx context.Context, take int, filter, limitFilter ServiceScope) (*dto.PageUserResponseType, error)
	Create(userDto *dto.CreateUserDto) (*response.ID, error)
	Update(userDto *dto.UpdateUserDto, scopes ...ServiceScope) error
	ChangeVisibility(scopes ...ServiceScope) error
	Delete(scopes ...ServiceScope) error
}

type userService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) UserService {
	return &userService{db}
}

func (s *userService) ModelWithContext(ctx context.Context) *gorm.DB {
	return s.db.WithContext(ctx).Model(&model.UserModel{})
}

func (s *userService) Model() *gorm.DB {
	return s.db.Model(&model.UserModel{})
}

func (s *userService) FindOne(ctx context.Context, scopes ...ServiceScope) (*model.UserModel, error) {

	var user model.UserModel
	{
		err := s.ModelWithContext(ctx).
			Scopes(scopes...).
			First(&user).
			Error

		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (s *userService) Find(ctx context.Context, scopes ...ServiceScope) ([]model.UserModel, error) {

	var moreModels []model.UserModel
	{
		err := s.ModelWithContext(ctx).
			Scopes(scopes...).
			Find(&moreModels).
			Error

		if err != nil {
			return nil, err
		}
	}

	return moreModels, nil
}

func (s *userService) Page(ctx context.Context, take int, filter, limitFilter ServiceScope) (*dto.PageUserResponseType, error) {

	tx := s.ModelWithContext(ctx)

	var total int64
	{
		txTotal := tx.Scopes(filter).Count(&total)
		if err := txTotal.Error; err != nil {
			return nil, err
		}
	}

	var users []*model.UserModel
	{
		if err := tx.Scopes(filter, limitFilter).
			Find(&users).Error; err != nil {
			return nil, err
		}
	}

	totalPages := int64(math.Ceil(float64(total) / float64(take)))

	return response.NewPaginateResponse(totalPages, users), nil
}

func (s *userService) Create(userDto *dto.CreateUserDto) (*response.ID, error) {

	user := model.UserModel{
		FullName:         userDto.FullName,
		Username:         userDto.Username,
		TelegramId:       userDto.TelegramId,
		RegTime:          userDto.RegTime,
		Phone:            userDto.Phone,
		Status:           userDto.Status,
		SubscriptionDate: userDto.SubscriptionDate,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &response.ID{Id: user.Id}, nil
}

func (s *userService) Update(userDto *dto.UpdateUserDto, scopes ...ServiceScope) error {
	return s.Model().Scopes(scopes...).Updates(userDto).Error
}

func (s *userService) Delete(scopes ...ServiceScope) error {
	return s.Model().Scopes(scopes...).Delete(nil).Error
}

func (s *userService) ChangeVisibility(scopes ...ServiceScope) error {
	var user model.UserModel

	if err := s.Model().Scopes(scopes...).First(&user).Error; err != nil {
		return err
	}

	newVisibility := !user.IsVisible

	if err := s.Model().Scopes(scopes...).Updates(map[string]interface{}{"is_visible": newVisibility}).Error; err != nil {
		return err
	}

	return nil
}
