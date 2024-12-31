package postgres

import (
	"context"
	"database/sql"

	"github.com/arthur-fontaine/culty/services/api/gen/db"
	"github.com/arthur-fontaine/culty/services/api/internal/domain/media"
	"github.com/stephenafamo/bob"
)

type MediaRepository struct {
	dbConn bob.DB
}

func NewMediaRepository(db *sql.DB) media.Repository {
	return MediaRepository{
		dbConn: bob.NewDB(db),
	}
}

func (m MediaRepository) GetMediumByID(ctx context.Context, id string) (*media.Media, error) {
	medium, err := db.FindMedium(ctx, m.dbConn, id)
	if err != nil {
		return nil, err
	}

	return &media.Media{
		ID:   medium.ID,
		Name: medium.Title,
	}, nil
}

func (m MediaRepository) ListMedia(ctx context.Context) ([]*media.Media, error) {
	allMedia, err := db.Media.Query().All(ctx, m.dbConn)
	if err != nil {
		return nil, err
	}

	var result []*media.Media
	for _, m := range allMedia {
		result = append(result, &media.Media{
			ID:   m.ID,
			Name: m.Title,
		})
	}

	return result, nil
}
