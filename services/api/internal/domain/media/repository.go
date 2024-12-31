package media

import "context"

type Repository interface {
	GetMediumByID(ctx context.Context, id string) (*Media, error)
	ListMedia(ctx context.Context) ([]*Media, error)
}
