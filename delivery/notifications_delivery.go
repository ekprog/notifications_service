package delivery

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"microservice_clean_design/app"
	"microservice_clean_design/app/core"
	"microservice_clean_design/domain"
	pb "microservice_clean_design/pkg/pb/api"
)

type NotificationsDeliveryService struct {
	pb.UnsafeNotificationsServiceServer
	log   core.Logger
	ucase domain.NotificationInteractor
}

func NewNotificationsDeliveryService(log core.Logger, ucase domain.NotificationInteractor) *NotificationsDeliveryService {
	return &NotificationsDeliveryService{
		log:   log,
		ucase: ucase,
	}
}

func (d *NotificationsDeliveryService) Init() error {
	app.InitGRPCService(pb.RegisterNotificationsServiceServer, pb.NotificationsServiceServer(d))
	return nil
}

func (d *NotificationsDeliveryService) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {

	userId, err := app.ExtractRequestUserId(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cannot extract user_id from context")
	}

	uCaseRes, err := d.ucase.GetAllNotifications(userId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetResponse{
		Status: &pb.Status{
			Code:    uCaseRes.StatusCode,
			Message: uCaseRes.StatusCode,
		},
	}

	if uCaseRes.StatusCode == domain.Success {

		var items []*pb.Notification
		for _, item := range uCaseRes.Notifications {
			items = append(items, &pb.Notification{
				Id:        item.Id,
				Title:     item.Title,
				Text:      item.Text,
				CreatedAt: timestamppb.New(item.CreatedAt),
			})
		}
		response.Notifications = items
	}

	return response, nil
}
