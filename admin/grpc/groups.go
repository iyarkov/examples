package grpc

import (
	"context"
	"examples/admin/database"
	"examples/admin/generated"
	"examples/admin/model"
	"github.com/google/uuid"
	"github.com/iyarkov/foundation/auth"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type groupsServer struct {
	generated.UnimplementedGroupsServer

	groupDAO database.GroupDAO
}

func (m *groupsServer) Get(ctx context.Context, request *generated.GetRequest) (*generated.Group, error) {
	panic("not implemented")
}

func (m *groupsServer) List(ctx context.Context, request *generated.GroupListRequest) (*generated.GroupListResponse, error) {
	panic("not implemented")
}

func (m *groupsServer) Create(ctx context.Context, request *generated.GroupCreateRequest) (*generated.GroupModificationResponse, error) {
	log := zerolog.Ctx(ctx)
	authToken := auth.AuthToken(ctx)
	if !authToken.IsInRole(auth.Admin) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	log.Debug().Msgf("create group[%s]", request.Name)
	now := time.Now()
	group := model.Group{
		Id:        uuid.New(),
		Name:      request.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := m.groupDAO.Create(ctx, &group); err != nil {
		log.Error().Err(err).Msg("create group failed")
		return nil, err
	}
	log.Debug().Msgf("group[%s] created, uuid: %v", request.Name, group.Id)
	response := generated.GroupModificationResponse{
		Code: generated.GroupModificationResponse_Ok,
		Result: &generated.Group{
			Id:        group.Id.String(),
			CreatedAt: uint64(group.CreatedAt.UnixMilli()),
			UpdatedAt: uint64(group.UpdatedAt.UnixMilli()),
			Name:      group.Name,
		},
	}
	return &response, nil
}

func (m *groupsServer) Update(ctx context.Context, request *generated.GroupUpdateRequest) (*generated.GroupModificationResponse, error) {
	panic("not implemented")
}
