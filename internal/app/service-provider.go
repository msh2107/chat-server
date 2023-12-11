package app

import (
	"context"
	"log"

	chatImpl "github.com/msh2107/chat-server/internal/api/chat-server"
	"github.com/msh2107/chat-server/internal/client/db"
	"github.com/msh2107/chat-server/internal/client/db/pg"
	"github.com/msh2107/chat-server/internal/client/db/transaction"
	chatRepository "github.com/msh2107/chat-server/internal/repository/chat"
	"github.com/msh2107/chat-server/internal/repository/message"
	"github.com/msh2107/chat-server/internal/repository/user"
	chatService "github.com/msh2107/chat-server/internal/service/chat-server"

	"github.com/msh2107/chat-server/internal/closer"
	"github.com/msh2107/chat-server/internal/config"
	"github.com/msh2107/chat-server/internal/config/env"
	"github.com/msh2107/chat-server/internal/repository"
	"github.com/msh2107/chat-server/internal/service"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient          db.Client
	txManager         db.TxManager
	chatRepository    repository.ChatRepository
	userRepository    repository.UserRepository
	messageRepository repository.MessageRepository

	chatService service.ChatService

	chatImpl *chatImpl.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = user.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = message.NewRepository(s.DBClient(ctx))
	}

	return s.messageRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewChatService(s.ChatRepository(ctx), s.MessageRepository(ctx), s.UserRepository(ctx), s.TxManager(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chatImpl.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chatImpl.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
