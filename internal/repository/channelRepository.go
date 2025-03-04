package repository

import (
	"database/sql"
	"fmt"

	"github.com/Gergenus/pkg"
)

type ChannelRepository interface {
	CreateChannel(name string, creatorId int) (int, error)
	DeleteChannel(channelId int) (int, error)
	GetOwner(channelId int) (int, error)
}

type PostgresChannelRepository struct {
	DB pkg.PostgresDatabase
}

func NewPostgresChannelRepository(db pkg.PostgresDatabase) PostgresChannelRepository {
	return PostgresChannelRepository{DB: db}
}

func (p *PostgresChannelRepository) CreateChannel(name string, creatorId int) (int, error) {
	var id int
	row := p.DB.GetDB().QueryRow("INSERT INTO channels (name, created_by) VALUES($1, $2) RETURNING id", name, creatorId)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *PostgresChannelRepository) DeleteChannel(channelId int) (int, error) {
	var id int
	row := p.DB.GetDB().QueryRow("DELETE FROM channels WHERE id=$1 RETURNING id", channelId)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *PostgresChannelRepository) GetOwner(channelId int) (int, error) {
	var id int
	row := p.DB.GetDB().QueryRow("SELECT created_by FROM channels WHERE id=$1", channelId)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("channel not found")
		}
		return 0, err
	}
	return id, nil
}
