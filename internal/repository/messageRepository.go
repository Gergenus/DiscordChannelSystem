package repository

import (
	"time"

	"github.com/Gergenus/internal/domain"
	"github.com/Gergenus/pkg"
)

type MessageRepository interface {
	CreateMessage(text string, channelId int, userId int) (int, error)
	DeleteMessage(messageId int) (int, error)
	ListMessages(channelId int, constraint int) ([]domain.Message, error)
	RetrieveMessagesDue(channelId int, constraint int, time time.Time) ([]domain.Message, error)
}

type PostgresMessageRepository struct {
	DB pkg.PostgresDatabase
}

func NewPostgresMessageRepository(db pkg.PostgresDatabase) PostgresMessageRepository {
	return PostgresMessageRepository{DB: db}
}

func (p *PostgresMessageRepository) CreateMessage(text string, channelId int, userId int) (int, error) {
	var id int
	row := p.DB.GetDB().QueryRow("INSERT INTO messages (content, channel_id, user_id) VALUES($1, $2, $3) RETURNING id", text, channelId, userId)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *PostgresMessageRepository) DeleteMessage(messageId int) (int, error) {
	var id int
	row := p.DB.GetDB().QueryRow("DELETE FROM messages WHERE id=$1 RETURNING id", messageId)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *PostgresMessageRepository) ListMessages(channelId int, constraint int) ([]domain.Message, error) {
	row, err := p.DB.GetDB().Query("SELECT * FROM messages WHERE channel_id=$1 ORDER BY created_at DESC LIMIT $2", channelId, constraint)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var messages []domain.Message

	for row.Next() {
		var message domain.Message

		err = row.Scan(&message.Id, &message.Content, &message.ChannelId, &message.UserId, &message.CreatedAt)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (p *PostgresMessageRepository) RetrieveMessagesDue(channelId int, constraint int, time time.Time) ([]domain.Message, error) {
	query := `
			SELECT id, content, channel_id, user_id, created_at
			FROM messages
			WHERE channel_id = $1 AND created_at > $2
			ORDER BY created_at DESC
			LIMIT $3
		    `
	rows, err := p.DB.GetDB().Query(query, channelId, time, constraint)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []domain.Message

	for rows.Next() {
		var message domain.Message
		err = rows.Scan(&message.Id, &message.Content, &message.ChannelId, &message.UserId, &message.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil

}
