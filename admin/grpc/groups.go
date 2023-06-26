package grpc

import (
	"context"
	"examples/admin/database"
	"examples/admin/generated"
	"examples/admin/model"
	"github.com/iyarkov/foundation/auth"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"unicode/utf8"
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
	//ctx, span := telemetry.Tracer.Start(ctx, "groupsServer#Create")
	//defer span.End()
	//span.SetAttributes(attribute.String("contextId", support.ContextId(ctx)))

	log := zerolog.Ctx(ctx)
	authToken := auth.AuthToken(ctx)
	if !authToken.IsInRole(auth.Admin) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	log.Debug().Msgf("create group[%s]", request.Name)

	if ok, messages := groupValid(request); !ok {
		log.Debug().Msgf("invalid request: %v", messages)
		grpcStatus := status.New(codes.InvalidArgument, "invalid request")
		details := generated.Messages{
			Message: messages,
		}
		grpcStatusWithDetails, err := grpcStatus.WithDetails(&details)
		if err != nil {
			log.Error().Err(err).Msg("failed to add details")
			return nil, grpcStatus.Err()
		}
		return nil, grpcStatusWithDetails.Err()
	}

	now := time.Now()
	group := model.Group{
		Name:      request.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := m.groupDAO.Create(ctx, &group); err == nil {
		log.Debug().Msgf("group[%s] created, id: %v", request.Name, group.Id)
		response := generated.GroupModificationResponse{
			Status: &generated.Status{
				Code: 0,
			},
			Result: &generated.Group{
				Id:        group.Id,
				CreatedAt: uint64(group.CreatedAt.UnixMilli()),
				UpdatedAt: uint64(group.UpdatedAt.UnixMilli()),
				Name:      group.Name,
			},
		}
		return &response, nil
	} else if err == database.ErrDuplicateName {
		log.Debug().Msgf("group[%s] created failed, name is not unique", request.Name)
		response := generated.GroupModificationResponse{
			Status: &generated.Status{
				Code:    uint32(generated.GroupModificationResponse_NameNonUnique.Number()),
				Details: []string{"Group with this name already exist"},
			},
		}
		return &response, nil
	} else {
		log.Error().Err(err).Msg("create group failed")
		return nil, err
	}
}

func (m *groupsServer) Update(ctx context.Context, request *generated.GroupUpdateRequest) (*generated.GroupModificationResponse, error) {
	panic("not implemented")
}

func groupValid(request *generated.GroupCreateRequest) (bool, []string) {
	messages := make([]string, 0)
	if request.Name == "" {
		messages = append(messages, "name is required")
	}
	length := utf8.RuneCountInString(request.Name)
	if length < 3 {
		messages = append(messages, "minimal name length is 3")
	}
	if length > 255 {
		messages = append(messages, "maximum name length is 255")
	}

	if len(messages) == 0 {
		return true, nil
	} else {
		return false, messages
	}
}
