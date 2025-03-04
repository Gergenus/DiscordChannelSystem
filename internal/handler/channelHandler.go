package handler

import (
	"errors"

	"github.com/Gergenus/internal/models"
	"github.com/Gergenus/internal/service"
	"github.com/labstack/echo/v4"
)

type ChannelHandler interface {
	CreateChannel(e echo.Context) error
	DeleteChannel(e echo.Context) error
}

type ChannelHttpHandler struct {
	channelService service.ChannelService
}

func NewChannelHttpHandler(channelService service.ChannelService) ChannelHttpHandler {
	return ChannelHttpHandler{channelService: channelService}
}

func (c *ChannelHttpHandler) CreateChannel(e echo.Context) error {
	var input models.ChannelReq
	err := e.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(400, "Bad request")
	}
	id, err := c.channelService.CreateChannel(input.Name, input.Uid)
	if err != nil {
		return echo.NewHTTPError(500, "Internal Error")
	}
	return e.JSON(200, id)
}

func (c *ChannelHttpHandler) DeleteChannel(e echo.Context) error {
	var input models.ChannelDel
	err := e.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(400, "Bad request")
	}
	id, err := c.channelService.DeleteChannel(input.Cid, input.Uid)
	if err != nil {
		if errors.Is(err, service.InvalidDeletion) {
			return echo.NewHTTPError(400, "Bad request")
		}
		return echo.NewHTTPError(500, "Internal Error")
	}
	return e.JSON(200, id)
}
