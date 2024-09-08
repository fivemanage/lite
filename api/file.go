package api

import "mime/multipart"

type UploadFile struct {
	FileHeader *multipart.FileHeader
	File       multipart.File
}
