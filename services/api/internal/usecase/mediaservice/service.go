package mediaservice

import (
	"github.com/arthur-fontaine/culty/services/api/internal/domain/media"
)

type MediaService struct {
	repo media.Repository
}

func NewMediaService(repo media.Repository) *MediaService {
	return &MediaService{
		repo: repo,
	}
}
