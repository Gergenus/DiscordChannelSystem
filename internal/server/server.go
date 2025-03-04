package server

import (
	"fmt"

	"github.com/Gergenus/config"
	"github.com/Gergenus/internal/handler"
	"github.com/Gergenus/internal/repository"
	"github.com/Gergenus/internal/service"
	"github.com/Gergenus/pkg"
	hasherpkg "github.com/Gergenus/pkg/Hasher"
	"github.com/Gergenus/pkg/tokens"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type Server interface {
	Start()
}

type EchoServer struct {
	app  *echo.Echo
	db   pkg.PostgresDatabase
	conf *config.Config
}

func NewEchoServer(db pkg.PostgresDatabase, conf *config.Config) EchoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)
	return EchoServer{
		app:  echoApp,
		db:   db,
		conf: conf,
	}
}

func (e *EchoServer) Start() {
	e.app.Use(middleware.Recover())
	e.app.Use(middleware.Logger())

	serv := fmt.Sprintf(":%d", e.conf.Server.Port)
	e.app.Logger.Fatal(e.app.Start(serv))
}

func (e *EchoServer) InitializationRouts() {

	userRepo := repository.NewPostgresUserRepository(e.db)
	channelRepo := repository.NewPostgresChannelRepository(e.db)
	tokenManager := tokens.NewJWTTokenManager(userRepo)
	newhasher := hasherpkg.NewCryptHasher()
	serviceauth := service.NewJWTauth(&newhasher, userRepo, tokenManager)
	channelService := service.NewPostgresChannelService(&channelRepo)
	echoauth := handler.NewEchoAuthHandler(serviceauth)
	echoChannel := handler.NewChannelHttpHandler(&channelService)
	echoMiddle := handler.NewEchoMiddleware(tokenManager)

	repoMessage := repository.NewPostgresMessageRepository(e.db)
	serviceMessage := service.NewPostgresMessageService(&repoMessage)
	echoMessages := handler.NewMessageHandler(serviceMessage)
	auth := e.app.Group("/auth")
	{
		auth.POST("/sign-up", echoauth.SignUp)
		auth.POST("/sign-in", echoauth.SignIn)
	}
	channel := e.app.Group("/channel")
	{
		channel.POST("/", echoChannel.CreateChannel)
		channel.DELETE("/", echoChannel.DeleteChannel)
	}
	messages := e.app.Group("/messages", echoMiddle.UserIndentity)
	{
		messages.POST("/", echoMessages.CreateMessage)
		messages.DELETE("/", echoMessages.DeleteMessage)
		messages.GET("/:channel_id", echoMessages.ListMessages)
	}
}
