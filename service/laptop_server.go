package service

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	pb "github.com/grpc-golang/pcbook/pb/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	Store LaptopStore
	
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{store}
}

func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()

	log.Printf(("received laptop request"))

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)

		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Id not valid")
		}
	} else {
		id, err := uuid.NewUUID()

		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate")
		}
		laptop.Id = id.String()

	}

	if ctx.Err() == context.DeadlineExceeded{
		log.Printf("deadline exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	if ctx.Err() == context.Canceled{
		log.Printf("cancelled")
		return nil, status.Error(codes.Canceled, "job is cancelled")
	}
	//save to inmemory

	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal

		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop to store %v", err)
	}

	log.Printf("save with id %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return res, nil

}

