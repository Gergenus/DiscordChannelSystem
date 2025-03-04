package service

import (
	"errors"

	"github.com/Gergenus/internal/repository"
)

var (
	InvalidDeletion = errors.New("You are not an owner of the channel")
)

type ChannelService interface {
	CreateChannel(name string, creatorId int) (int, error)
	DeleteChannel(channelId int, userId int) (int, error)
}

type PostgresChannelService struct {
	channelRepo repository.ChannelRepository
}

func NewPostgresChannelService(channelRepo repository.ChannelRepository) PostgresChannelService {
	return PostgresChannelService{channelRepo: channelRepo}
}

func (p *PostgresChannelService) CreateChannel(name string, creatorId int) (int, error) {
	id, err := p.channelRepo.CreateChannel(name, creatorId)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *PostgresChannelService) DeleteChannel(channelId int, userId int) (int, error) {
	ownerId, err := p.channelRepo.GetOwner(channelId)
	if err != nil {
		return 0, err
	}
	if ownerId != userId {
		return 0, InvalidDeletion
	}
	id, err := p.channelRepo.DeleteChannel(channelId)
	return id, nil
}
