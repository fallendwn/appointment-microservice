package client

import (
	"context"

	pb "github.com/fallendwn/appointment-proto/doctor"
	"github.com/fallendwn/appointment/appointment-service/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type DoctorClient struct {
	client pb.DoctorServiceClient
}

func NewDoctorClient(address string) (*DoctorClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &DoctorClient{client: pb.NewDoctorServiceClient(conn)}, nil
}

func (c *DoctorClient) CheckDoctorExists(ctx context.Context, doctorID string) (bool, error) {
	_, err := c.client.GetDoctor(ctx, &pb.GetDoctorRequest{Id: doctorID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				return false, model.ErrDoctorNotFound
			}
		}
		return false, model.ErrDoctorUnavailable
	}
	return true, nil
}
