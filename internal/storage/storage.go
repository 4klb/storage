package storage

import "github.com/4klb/storage_ozon/internal/file"

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(filename string, blob []byte) *Storage {
	file.File{}
	return &Storage{}
}
