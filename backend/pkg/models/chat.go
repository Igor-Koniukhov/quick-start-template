package models

import (
	"database/sql"
	"time"
)

const UsersTable = "users"
const ImagesTable = "images"
const ChatsTable = "chats"
const MessagesTable = "messages"
const ChatParticipantsTable = "chat_participants"
const ChatImagesTable = "chat_images"

type User struct {
	ID             int       `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	ProfileImageId int       `json:"profile_image_id"`
	Password       string    `json:"password"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Chat struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	AdminID   int       `json:"admin_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ChatParticipant struct {
	UserID int `json:"user_id"`
	ChatID int `json:"chat_id"`
}

type Message struct {
	ID        int       `json:"id"`
	ChatID    int       `json:"chat_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	ImagePath string    `json:"image_path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Image struct {
	ID        int          `json:"id"`
	ChatID    int          `json:"chat_id"`
	UserID    int          `json:"user_id"`
	Context   string       `json:"context"`
	IsRead    int          `json:"is_read"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
