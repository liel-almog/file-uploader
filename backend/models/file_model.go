package models

import "time"

type UploadFileMetadateDTO struct {
	FileName string `json:"fileName" binding:"required" validate:"min=1"`
	Size     int64  `json:"size" binding:"required" validate:"min=1"`
	MimeType string `json:"mimeType" binding:"required" validate:"min=1"`
}

type FileMetadata struct {
	// embed the UploadFileMetadateDTO
	UploadFileMetadateDTO
	UserID int64 `json:"userId" binding:"required" validate:"min=1"`
}

type FileResDTO struct {
	FileID      int64     `json:"fileId"`
	DisplayName string    `json:"displayName"`
	UploadedAt  time.Time `json:"uploadedAt" db:"uploaded_at"`
	Favorite    bool      `json:"favorite"`
	Size        int64     `json:"size"`
	Status      string    `json:"status"`
}