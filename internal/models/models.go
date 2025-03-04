package models

type ChannelReq struct {
	Uid  int    `json:"uid"`
	Name string `json:"name"`
}

type ChannelDel struct {
	Uid int `json:"uid"`
	Cid int `json:"cid"`
}

type AuthInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type MessageInput struct {
	Text      string `json:"text"`
	ChannelId int    `json:"channel_id"`
	UserId    int    `json:"user_id"`
}

type DeleteMessageInput struct {
	MessageId int `json:"message_id"`
}

type RetrieveMessageInput struct {
	ChannelId  int `json:"channel_id"`
	Constraint int `json:"constraint"`
}
