package mediaservice

import (
	"context"
	"log"

	"connectrpc.com/connect"
	mediav1 "github.com/arthur-fontaine/culty/services/api/gen/proto/media/v1"
)

func (s *MediaService) GetMedium(
	ctx context.Context,
	req *connect.Request[mediav1.MediumRequest],
) (*connect.Response[mediav1.Medium], error) {
	log.Println("Request headers: ", req.Header())

	media, err := s.repo.GetMediumByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&mediav1.Medium{
		Id:    media.ID,
		Title: media.Name,
	})
	res.Header().Set("Media-Version", "v1")

	return res, nil
}
