package domain

import "time"

type User struct {
	Uid          int
	UserName     string
	HashPassword string
}

type Channel struct {
	Id        int
	Name      string
	CreatorId int
}

type Message struct {
	Id        int
	Content   string
	ChannelId int
	UserId    int
	CreatedAt time.Time
}
