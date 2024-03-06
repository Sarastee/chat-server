package app

import (
	"context"
	"log"

	"github.com/sarastee/chat-server/internal/client/db"
	"github.com/sarastee/chat-server/internal/client/db/pg"

	"github.com/sarastee/chat-server/internal/api/chat"
	"github.com/sarastee/chat-server/internal/closer"
	"github.com/sarastee/chat-server/internal/config"
	"github.com/sarastee/chat-server/internal/config/env"
	"github.com/sarastee/chat-server/internal/repository"
	chatRepository "github.com/sarastee/chat-server/internal/repository/chat"
	messageRepository "github.com/sarastee/chat-server/internal/repository/message"
	userRepository "github.com/sarastee/chat-server/internal/repository/user"
	"github.com/sarastee/chat-server/internal/service"
	chatService "github.com/sarastee/chat-server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   *config.PgConfig
	grpcConfig *config.GRPCConfig

	dbClient  db.Client
	txManager db.TxManager

	chatRepo    repository.ChatRepository
	userRepo    repository.UserRepository
	messageRepo repository.MessageRepository

	chatService service.ChatService

	chatImpl *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PgConfig ..
func (s *serviceProvider) PgConfig() *config.PgConfig {
	if s.pgConfig == nil {
		cfgSearcher := env.NewPgCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get PG config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

// GRPCConfig ..
func (s *serviceProvider) GRPCConfig() *config.GRPCConfig {
	if s.grpcConfig == nil {
		cfgSearcher := env.NewGRPCCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get PG config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

// DBClient ..
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PgConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create DB Client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping DB: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

// TxManager ..
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = pg.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// ChatRepository ..
func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepo == nil {
		s.chatRepo = chatRepository.NewRepo(s.DBClient(ctx))
	}

	return s.chatRepo
}

// UserRepository ..
func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepo == nil {
		s.userRepo = userRepository.NewRepo(s.DBClient(ctx))
	}

	return s.userRepo
}

// MessageRepository ..
func (s *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepository {
	if s.messageRepo == nil {
		s.messageRepo = messageRepository.NewRepo(s.DBClient(ctx))
	}

	return s.messageRepo
}

// ChatService ..
func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.TxManager(ctx),
			s.ChatRepository(ctx),
			s.UserRepository(ctx),
			s.MessageRepository(ctx),
		)
	}

	return s.chatService
}

// ChatImpl ..
func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
