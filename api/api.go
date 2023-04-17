package api

import (
	"context"
	"encoding/json"
	"flag"
	"net"
	"net/http"
	"time"

	"github.com/jordation/layermon/db"
	"github.com/jordation/layermon/stats"
	pb_stats "github.com/jordation/layermon/stats/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type API struct {
	statsClient pb_stats.StatsServiceClient
}

const (
	port = ":9000"
)

var (
	GrpcPort = *flag.String("port", port, "GRPC API port")
	GrpcAddr = *flag.String("grpc addr", "localhost"+port, "grpc address")
)

func StartGrpcServer() {
	lis, err := net.Listen("tcp", GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
	db := db.GetNewDbConnection()
	statServer := stats.NewStatsRepo(db)
	grpcServer := grpc.NewServer()
	pb_stats.RegisterStatsServiceServer(grpcServer, statServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
	log.Println("starting grpc server on ", lis.Addr())
}

func GetApi() *API {
	flag.Parse()
	return &API{
		statsClient: stats.GetStatsClient(GrpcPort),
	}
}

func (A *API) HandleListStats(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	var q pb_stats.GetStatByIdRequest
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		log.Info(err)
		return
	}
	res, err := A.statsClient.GetStatById(ctx, &q)
	if err != nil {
		log.Info(err)
		return
	}
	json.NewEncoder(w).Encode(res.GetMessage())
}
