package storage

import (
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/alfianyulianto/go-room-managament/halpers"
)

type FileStorage interface {
	SaveFile(file *multipart.FileHeader, subFolder string) (string, error)
}

type LocalFileStorage struct {
	BasePath string
}

func NewLocalFileStorage() *LocalFileStorage {
	return &LocalFileStorage{BasePath: "./uploads"}
}
func (s LocalFileStorage) SaveFile(file *multipart.FileHeader, subFolder string) (string, error) {
	ext := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + ext

	//make folder
	dirPath := filepath.Join(s.BasePath, subFolder)
	err := os.MkdirAll(dirPath, os.ModePerm)
	halpers.IfPanicError(err)

	filePath := filepath.Join(dirPath, newFileName)

	src, err := file.Open()
	halpers.IfPanicError(err)

	fileDestination, err := os.Create(filePath)
	halpers.IfPanicError(err)
	defer fileDestination.Close()

	_, err = io.Copy(fileDestination, src)
	halpers.IfPanicError(err)

	return filepath.ToSlash(filePath), nil
}
