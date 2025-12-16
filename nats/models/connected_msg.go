package models

import "time"

type UserConnectedMsg struct {
	UserID    int64     `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
}

func (e UserConnectedMsg) GetUserID() int64        { return e.UserID }
func (e UserConnectedMsg) GetTimestamp() time.Time { return e.Timestamp }

type UserDisconnectedMsg struct {
	UserID    int64     `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
}

func (e UserDisconnectedMsg) GetUserID() int64        { return e.UserID }
func (e UserDisconnectedMsg) GetTimestamp() time.Time { return e.Timestamp }

type UserOnlineMsg struct {
	IsOnline bool `json:"is_online"`
}

type PresenceStatusEvent struct {
	UserID    int64     `json:"user_id"`
	IsOnline  bool      `json:"is_online"`
	Timestamp time.Time `json:"timestamp"`
}

type PresenceHeartbeatMsg struct {
	UserID    int64     `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
}

type MessageCreatedEvent struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	ChatID      string `json:"chat_id"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   *int64 `json:"updated_at"`
	MessageType string `json:"message_type"`
	Status      string `json:"status"`
	SenderID int64 `json:"sender_id"`
}
