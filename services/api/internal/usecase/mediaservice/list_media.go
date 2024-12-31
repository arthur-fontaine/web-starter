package mediaservice

import (
	"context"

	"connectrpc.com/connect"
	mediav1 "github.com/arthur-fontaine/culty/services/api/gen/proto/media/v1"
)

func (s *MediaService) ListMedia(
	ctx context.Context,
	req *connect.Request[mediav1.Empty],
) (*connect.Response[mediav1.MediumList], error) {
	media, err := s.repo.ListMedia(ctx)
	if err != nil {
		return nil, err
	}

	mediumList := &mediav1.MediumList{}
	for _, m := range media {
		mediumList.Media = append(mediumList.Media, &mediav1.Medium{
			Id:    m.ID,
			Title: m.Name,
		})
	}

	res := connect.NewResponse(mediumList)
	res.Header().Set("Media-Version", "v1")

	return res, nil
}
