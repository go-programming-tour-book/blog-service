package service

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/go-programming-tour-book/blog-service/pkg/upload"
)

type FileInfo struct {
	Name string
	Type upload.FileType
	Dst  string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName
	if !upload.CheckExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}
	if !upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}
	if !upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	if !upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	return &FileInfo{
		Name: fileName,
		Type: fileType,
		Dst:  dst,
	}, nil
}
