package service

import (
	"context"
	"math"
	"strings"

	"github.com/Mirsadikovv/tvgo/playlist/dto"
	"github.com/Mirsadikovv/tvgo/playlist/model"
	"github.com/fobus1289/ufa_shared/http/response"
	"gorm.io/gorm"
)

type ServiceScope = func(d *gorm.DB) *gorm.DB

type PlaylistService interface {
	FindOne(ctx context.Context, scopes ...ServiceScope) (*model.PlaylistModel, error)
	Find(ctx context.Context, scopes ...ServiceScope) ([]model.PlaylistModel, error)
	Page(ctx context.Context, take int, filter, limitFilter ServiceScope) (*dto.PagePlaylistResponseType, error)
	Create(playlistDto *dto.CreatePlaylistDto) (*response.ID, error)
	Update(playlistDto *dto.UpdatePlaylistDto, scopes ...ServiceScope) error
	ChangeVisibility(scopes ...ServiceScope) error
	Delete(scopes ...ServiceScope) error
}

type playlistService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) PlaylistService {
	return &playlistService{db}
}

func (s *playlistService) ModelWithContext(ctx context.Context) *gorm.DB {
	return s.db.WithContext(ctx).Model(&model.PlaylistModel{})
}

func (s *playlistService) Model() *gorm.DB {
	return s.db.Model(&model.PlaylistModel{})
}

func (s *playlistService) FindOne(ctx context.Context, scopes ...ServiceScope) (*model.PlaylistModel, error) {

	var playlist model.PlaylistModel
	{
		err := s.ModelWithContext(ctx).
			Scopes(scopes...).
			First(&playlist).
			Error

		if err != nil {
			return nil, err
		}
	}

	return &playlist, nil
}

func (s *playlistService) Find(ctx context.Context, scopes ...ServiceScope) ([]model.PlaylistModel, error) {

	var moreModels []model.PlaylistModel
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

func (s *playlistService) Page(ctx context.Context, take int, filter, limitFilter ServiceScope) (*dto.PagePlaylistResponseType, error) {

	tx := s.ModelWithContext(ctx)

	var total int64
	{
		txTotal := tx.Scopes(filter).Count(&total)
		if err := txTotal.Error; err != nil {
			return nil, err
		}
	}

	var playlists []*model.PlaylistModel
	{
		if err := tx.Scopes(filter, limitFilter).
			Find(&playlists).Error; err != nil {
			return nil, err
		}
	}

	totalPages := int64(math.Ceil(float64(total) / float64(take)))

	return response.NewPaginateResponse(totalPages, playlists), nil
}

func (s *playlistService) Create(playlistDto *dto.CreatePlaylistDto) (*response.ID, error) {

	playlist := model.PlaylistModel{
		Name:        playlistDto.Name,
		Description: playlistDto.Description,
		LogoLink:    playlistDto.LogoLink,
		ChannelLink: playlistDto.ChannelLink,
		Type:        strings.ToUpper(string(playlistDto.Type)),
	}

	if err := s.db.Create(&playlist).Error; err != nil {
		return nil, err
	}

	return &response.ID{Id: playlist.Id}, nil
}

func (s *playlistService) Update(playlistDto *dto.UpdatePlaylistDto, scopes ...ServiceScope) error {
	return s.Model().Scopes(scopes...).Updates(playlistDto).Error
}

func (s *playlistService) Delete(scopes ...ServiceScope) error {
	return s.Model().Scopes(scopes...).Delete(nil).Error
}

func (s *playlistService) ChangeVisibility(scopes ...ServiceScope) error {
	var playlist model.PlaylistModel

	if err := s.Model().Scopes(scopes...).First(&playlist).Error; err != nil {
		return err
	}

	newVisibility := !playlist.IsVisible

	if err := s.Model().Scopes(scopes...).Updates(map[string]interface{}{"is_visible": newVisibility}).Error; err != nil {
		return err
	}

	return nil
}
