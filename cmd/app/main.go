package main

import (
	"github.com/miladvatankhah/maker-checker/configs"
	"github.com/miladvatankhah/maker-checker/pkg/clients/rabbit"
	"log"

	"github.com/miladvatankhah/maker-checker/internal/message_approval/application/use_cases"
	"github.com/miladvatankhah/maker-checker/internal/message_approval/infrastructure/messaging/rabbitmq"
	"github.com/miladvatankhah/maker-checker/internal/message_approval/infrastructure/persistence/postgres"
	"github.com/miladvatankhah/maker-checker/internal/message_approval/infrastructure/transport/http"
	"github.com/miladvatankhah/maker-checker/internal/message_approval/presentation/http/v1"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
	postgresClient "github.com/miladvatankhah/maker-checker/pkg/clients/postgres"
)

func init() {
	if err := godotenv.Load(".env.dev"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load DB config: %v", err)
	}

	pgClient, err := postgresClient.NewPostgresClient(cfg.Postgres)
	if err != nil {
		log.Fatalf("Failed to initialize PostgreSQL client: %v", err)
	}
	defer pgClient.Close()

	// Connect to RabbitMQ
	r, err := rabbit.DialWithDefaults(cfg.Rabbit)
	//amqpConn, err := amqp.Dial(cfg.Rabbit.Url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer r.Shutdown()

	//amqpChannel, err := amqpConn.Channel()
	//if err != nil {
	//	log.Fatalf("Failed to open a channel: %v", err)
	//}
	//defer amqpChannel.Close()

	// Initialize repositories
	userRepo := postgres.NewUserRepositoryImpl(pgClient.DB)
	messageRepo := postgres.NewMessageRepositoryImpl(pgClient.DB)

	// Initialize RabbitMQ event publisher
	eventPublisher := events.NewRabbitMQEventPublisher(r.Chan())

	// Initialize use cases
	createMessageUseCase := use_cases.NewCreateMessageUseCase(userRepo, messageRepo)
	approveMessageUseCase := use_cases.NewApproveMessageUseCase(messageRepo, eventPublisher)
	rejectMessageUseCase := use_cases.NewRejectMessageUseCase(messageRepo)
	registerUserUseCase := use_cases.NewRegisterUserUseCase(userRepo)

	// Initialize HTTP handlers for different API versions
	messageHandlerV1 := v1.NewMessageHandler(createMessageUseCase, approveMessageUseCase, rejectMessageUseCase)
	userHandlerV1 := v1.NewUserHandler(registerUserUseCase)

	// Initialize and start HTTP server
	httpServer := http.NewHTTPServer(cfg.Server, messageHandlerV1, userHandlerV1)

	// Start the HTTP server
	if err := httpServer.Start(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
