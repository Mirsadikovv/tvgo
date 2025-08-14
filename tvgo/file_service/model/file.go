package model

import (
	"os"
	"path/filepath"
	"time"
)

type FileModel struct {
	Id        int64      `json:"id" gorm:"primaryKey"`
	Owner     string     `json:"owner" gorm:"not null"`
	Category  string     `json:"category" gorm:"not null"`
	Sign      string     `json:"sign"`
	DirName   string     `json:"dirName"`
	Filename  string     `json:"filename"`
	MimeType  string     `json:"mimeType"`
	Size      int64      `json:"size"`
	IsVisible bool       `json:"isVisible" gorm:"default:true"`
	CreatedAt *time.Time `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"autoUpdateTime:true"`
}

func (FileModel) TableName() string {
	return "files"
}

func (f *FileModel) BuildFilePath() string {
	return filepath.Join("./", f.DirName, f.Category, f.Owner, f.Filename)
}

func (f *FileModel) CanDelete() bool {
	file, err := os.Stat(f.BuildFilePath())

	if err != nil {
		return true
	}

	return !file.IsDir()
}

func (f *FileModel) Remove() error {
	return os.Remove(f.BuildFilePath())
}
