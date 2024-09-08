package storage

import "github.com/fivemanage/lite/internal/storage/s3"

type StorageLayer interface {
	UploadFile() error
	DeleteFile() error
}

func New(provider string) StorageLayer {
	switch provider {
	case "s3":
		return s3.New()
	}

	return nil
}
