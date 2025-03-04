package handler

import (
	"log"
	"strconv"

	"github.com/Gergenus/internal/models"
	"github.com/Gergenus/internal/service"
	"github.com/labstack/echo/v4"
)

type MessageHandlerInterface interface {
	CreateMessage(ctx echo.Context) error
	DeleteMessage(ctx echo.Context) error
	ListMessages(ctx echo.Context) error
}

type MessageHandler struct {
	MessageService service.MessageService
}

func NewMessageHandler(MessageService service.MessageService) MessageHandler {
	return MessageHandler{MessageService: MessageService}
}

func (m *MessageHandler) CreateMessage(ctx echo.Context) error {
	var input models.MessageInput

	err := ctx.Bind(&input)
	if err != nil {
		log.Println("CreateMessage handler err", err)
		return echo.NewHTTPError(400, Messenger{Message: "Bad req"})
	}

	id, err := m.MessageService.CreateMessage(input.Text, input.ChannelId, input.UserId)
	if err != nil {
		log.Println("CreateMessage handler err", err)
		return echo.NewHTTPError(500, Messenger{Message: "Internal error"})
	}
	return ctx.JSON(200, Messenger{Message: id})
}

func (m *MessageHandler) DeleteMessage(ctx echo.Context) error {
	var input models.DeleteMessageInput

	err := ctx.Bind(&input)
	if err != nil {
		log.Println("DeleteMessage handler err", err)
		return echo.NewHTTPError(400, Messenger{Message: "Bad req"})
	}
	id, err := m.MessageService.DeleteMessage(input.MessageId)
	if err != nil {
		log.Println("DeleteMessage handler err", err)
		return echo.NewHTTPError(500, Messenger{Message: "Internal error"})
	}
	return ctx.JSON(200, Messenger{Message: id})
}

func (m *MessageHandler) ListMessages(ctx echo.Context) error {

	channel_id := ctx.Param("channel_id")
	offset := ctx.QueryParam("offset")
	num_channel_id, err := strconv.Atoi(channel_id)
	if err != nil {
		log.Println("ListMessages handler atoi err", err)
		return echo.NewHTTPError(500, Messenger{Message: "Internal error"})
	}

	num_offset, err := strconv.Atoi(offset)
	if err != nil {
		log.Println("ListMessages handler atoi err", err)
		return echo.NewHTTPError(500, Messenger{Message: "Internal error"})
	}

	messages, err := m.MessageService.ListMessages(num_channel_id, num_offset)
	if err != nil {
		log.Println("ListMessages handler err", err)
		return echo.NewHTTPError(500, Messenger{Message: "Internal error"})
	}
	return ctx.JSON(200, messages)
}
