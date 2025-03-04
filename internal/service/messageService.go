package service

import (
	"time"

	"github.com/Gergenus/internal/domain"
	"github.com/Gergenus/internal/repository"
)

type MessageService interface {
	CreateMessage(text string, channelId int, userId int) (int, error)
	DeleteMessage(messageId int) (int, error)
	ListMessages(channelId int, constraint int) ([]domain.Message, error)
	RetrieveMessagesDue(channelId int, constraint int, time time.Time) ([]domain.Message, error)
}

type PostgresMessageService struct {
	messageRepo repository.MessageRepository
}

func NewPostgresMessageService(messageRepo repository.MessageRepository) *PostgresMessageService {
	return &PostgresMessageService{messageRepo: messageRepo}
}

func (p *PostgresMessageService) CreateMessage(text string, channelId int, userId int) (int, error) {
	id, err := p.messageRepo.CreateMessage(text, channelId, userId)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (p *PostgresMessageService) DeleteMessage(messageId int) (int, error) {
	id, err := p.messageRepo.DeleteMessage(messageId)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (p *PostgresMessageService) ListMessages(channelId int, constraint int) ([]domain.Message, error) {
	messages, err := p.messageRepo.ListMessages(channelId, constraint)
	if err != nil {
		return nil, err
	}
	return messages, err
}

func (p *PostgresMessageService) RetrieveMessagesDue(channelId int, constraint int, time time.Time) ([]domain.Message, error) {
	messages, err := p.messageRepo.RetrieveMessagesDue(channelId, constraint, time)
	if err != nil {
		return nil, err
	}
	return messages, err
}
