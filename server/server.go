package server

import (
	"context"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/kwanok/minsim-api/config"
	"github.com/kwanok/minsim-api/database"
	"github.com/kwanok/minsim-api/server/comment"
	"github.com/kwanok/minsim-api/server/minsim"
	"github.com/kwanok/minsim-api/server/post"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Server interface {
	Start() error
}

type server struct {
	grpcServer *grpc.Server
	config     *config.Config
}

func loggingMiddleware() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Printf("Request - Method: %s, Duration: %s", info.FullMethod, time.Since(time.Now()))
		return handler(ctx, req)
	}
}

func NewServer(
	c *config.Config,
) Server {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
			loggingMiddleware(),
		),
	)

	return &server{
		grpcServer: grpcServer,
		config:     c,
	}
}

func (s *server) Start() error {
	listen, err := net.Listen("tcp", ":8100")
	if err != nil {
		return err
	}

	db := database.NewDB(s.config)
	db.Migrate()

	post.Register(s.grpcServer, db)
	minsim.Register(s.grpcServer, db)
	comment.Register(s.grpcServer, db)

	log.Printf("Server listening on %s", listen.Addr().String())

	return s.grpcServer.Serve(listen)
}
