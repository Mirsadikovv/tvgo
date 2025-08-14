package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"time"

	"github.com/Mirsadikovv/tvgo/file_service/dto"
	"github.com/Mirsadikovv/tvgo/file_service/model"
	"github.com/fobus1289/ufa_shared/http/response"
	"gorm.io/gorm"
)

type ServiceScope = func(d *gorm.DB) *gorm.DB

type FileService interface {
	FindOne(ctx context.Context, scopes ...ServiceScope) (*model.FileModel, error)
	Find(ctx context.Context, scopes ...ServiceScope) ([]model.FileModel, error)
	Page(ctx context.Context, take int, filter, limitFilter ServiceScope) (*dto.PageFileResponseType, error)
	Create(fileDto *dto.CreateFileDto) (*response.ID, error)
	Update(fileDto *dto.UpdateFileDto, scopes ...ServiceScope) error
	Replace(fileDto *dto.ReplaceFileDto, scopes ...ServiceScope) error
	ChangeVisibility(scopes ...ServiceScope) error
	Delete(scopes ...ServiceScope) error
}

type fileService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) FileService {
	return &fileService{db}
}

func (s *fileService) ModelWithContext(ctx context.Context) *gorm.DB {
	return s.db.WithContext(ctx).Model(&model.FileModel{})
}

func (s *fileService) Model() *gorm.DB {
	return s.db.Model(&model.FileModel{})
}

func (s *fileService) FindOne(ctx context.Context, scopes ...ServiceScope) (*model.FileModel, error) {
	var file model.FileModel
	{
		err := s.ModelWithContext(ctx).
			Scopes(scopes...).
			First(&file).
			Error

		if err != nil {
			return nil, err
		}
	}

	return &file, nil
}

func (s *fileService) Find(ctx context.Context, scopes ...ServiceScope) ([]model.FileModel, error) {
	var moreModels []model.FileModel
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

func (s *fileService) Page(ctx context.Context, take int, filter, limitFilter ServiceScope) (*dto.PageFileResponseType, error) {
	tx := s.ModelWithContext(ctx)

	var total int64
	{
		txTotal := tx.Scopes(filter).Count(&total)
		if err := txTotal.Error; err != nil {
			return nil, err
		}
	}

	var files []*model.FileModel
	{
		if err := tx.Scopes(filter, limitFilter).
			Find(&files).Error; err != nil {
			return nil, err
		}
	}

	totalPages := int64(math.Ceil(float64(total) / float64(take)))

	return response.NewPaginateResponse(totalPages, files), nil
}

func (s *fileService) Create(fileDto *dto.CreateFileDto) (*response.ID, error) {
	const FILES = "files"

	var (
		fileHeader = fileDto.FileHeader
		category   = fileDto.Category
		owner      = fileDto.Owner
		sign       = fileDto.Sign
	)

	ownerDirname := path.Join("./", FILES, category, owner)
	{
		if err := os.MkdirAll(ownerDirname, 0755); err != nil {
			return nil, err
		}
	}

	ownerFilename := generateRandomFilename(fileHeader.Filename)
	ownerFullFilePath := path.Join(ownerDirname, ownerFilename)

	multipart, err := fileHeader.Open()
	{
		if err != nil {
			return nil, err
		}
		defer multipart.Close()
	}

	ownerFile, err := os.Create(ownerFullFilePath)
	{
		if err != nil {
			return nil, err
		}
		defer ownerFile.Close()
	}

	fileModel := model.FileModel{
		DirName:  FILES,
		Category: category,
		Filename: path.Base(ownerFilename),
		Sign:     sign,
		MimeType: path.Ext(fileHeader.Filename),
		Owner:    owner,
		Size:     fileHeader.Size,
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&fileModel).Error; err != nil {
			return err
		}

		if _, err := io.Copy(ownerFile, multipart); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		os.Remove(ownerFullFilePath)
		return nil, err
	}

	return &response.ID{Id: fileModel.Id}, nil
}

func (s *fileService) Update(fileDto *dto.UpdateFileDto, scopes ...ServiceScope) error {
	return s.Model().Scopes(scopes...).Updates(fileDto).Error
}

func (s *fileService) Replace(fileDto *dto.ReplaceFileDto, scopes ...ServiceScope) error {
	const FILES = "files"

	var existingFile model.FileModel
	{
		if err := s.Model().Scopes(scopes...).First(&existingFile).Error; err != nil {
			return err
		}
	}

	oldFilePath := existingFile.BuildFilePath()

	var (
		fileHeader = fileDto.FileHeader
		category   = existingFile.Category
		owner      = existingFile.Owner
	)

	ownerDirname := path.Join("./", FILES, category, owner)
	{
		if err := os.MkdirAll(ownerDirname, 0755); err != nil {
			return err
		}
	}

	ownerFilename := generateRandomFilename(fileHeader.Filename)
	ownerFullFilePath := path.Join(ownerDirname, ownerFilename)

	multipart, err := fileHeader.Open()
	{
		if err != nil {
			return err
		}
		defer multipart.Close()
	}

	ownerFile, err := os.Create(ownerFullFilePath)
	{
		if err != nil {
			return err
		}
		defer ownerFile.Close()
	}

	updateData := map[string]interface{}{
		"filename":  path.Base(ownerFilename),
		"mime_type": path.Ext(fileHeader.Filename),
		"size":      fileHeader.Size,
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.FileModel{}).Scopes(scopes...).Updates(updateData).Error; err != nil {
			return err
		}

		if _, err := io.Copy(ownerFile, multipart); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		os.Remove(ownerFullFilePath)
		return err
	}

	if oldFilePath != "" {
		os.Remove(oldFilePath)
	}

	return nil
}

func (s *fileService) Delete(scopes ...ServiceScope) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var fileModel model.FileModel
		{
			if err := tx.Model(&model.FileModel{}).Scopes(scopes...).First(&fileModel).Error; err != nil {
				return err
			}
		}

		if err := tx.Model(&model.FileModel{}).Scopes(scopes...).Delete(nil).Error; err != nil {
			return err
		}

		if !fileModel.CanDelete() {
			return errors.New("can not delete")
		}

		return fileModel.Remove()
	})
}

func (s *fileService) ChangeVisibility(scopes ...ServiceScope) error {
	var file model.FileModel

	if err := s.Model().Scopes(scopes...).First(&file).Error; err != nil {
		return err
	}

	newVisibility := !file.IsVisible

	if err := s.Model().Scopes(scopes...).Updates(map[string]interface{}{"is_visible": newVisibility}).Error; err != nil {
		return err
	}

	return nil
}

// generateRandomFilename generates a random filename with the original extension
func generateRandomFilename(originalFilename string) string {
	ext := path.Ext(originalFilename)
	// Generate a simple random filename using timestamp and random number
	// This is a basic implementation - you can enhance it as needed
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("file_%d%s", timestamp, ext)
}
