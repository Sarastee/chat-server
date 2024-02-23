package main

import (
	"context"
	"log"
	"net"

	"github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sarastee/chat-server/internal/config/env"
	desc "github.com/sarastee/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	desc.UnimplementedChatV1Server
	pool *pgxpool.Pool
	sq   squirrel.StatementBuilderType
}

func main() {
	ctx := context.Background()

	cfg, err := env.New()
	if err != nil {
		log.Fatalf("cannot get config: %v", err)
	}

	lis, err := net.Listen(cfg.GRPC.Protocol, cfg.GRPC.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, cfg.Postgres.DSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{
		pool: pool,
	})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("chat create request")

	builderInsert := s.sq.Insert("chats").
		PlaceholderFormat(squirrel.Dollar).
		Columns("users").
		Values(req.GetUsernames()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}
	var chatID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		log.Fatalf("failed to insert chat: %v", err)
	}

	log.Printf("inserted chat with id: %d", chatID)

	return &desc.CreateResponse{
		Id: chatID,
	}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("chat delete request")

	builderDelete := s.sq.Delete("chats").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": req.GetId()})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Fatalf("failed to build query to delete chat: %v", err)
	}

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to execute query to delete chat: %v", err)
	}

	log.Printf("chat was deleted")
	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(_ context.Context, _ *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("send message request")

	log.Printf("Message: %s", gofakeit.BeerName())

	return &emptypb.Empty{}, nil
}
