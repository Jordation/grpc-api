package stats

import (
	"context"
	"fmt"
	"log"
	"sync"

	pb "github.com/jordation/layermon/stats/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	once   sync.Once
	client pb.StatsServiceClient
)

type StatsRepo struct {
	pb.UnimplementedStatsServiceServer
	db *map[int32]string
}

func NewStatsRepo(db *map[int32]string) pb.StatsServiceServer {
	return &StatsRepo{db: db}
}

func GetStatsClient(GrpcPort string) pb.StatsServiceClient {
	once.Do(func() {
		conn, err := grpc.Dial(GrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		client = pb.NewStatsServiceClient(conn)
	})
	return client
}

func (s *StatsRepo) GetStatById(ctx context.Context, in *pb.GetStatByIdRequest) (*pb.GetStatByIdResponse, error) {
	v, ok := (*s.db)[in.GetId()]
	if ok {
		return &pb.GetStatByIdResponse{Message: v}, nil
	} else {
		return nil, fmt.Errorf("id not found")
	}
}
