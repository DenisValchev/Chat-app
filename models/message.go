package models

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	Body       string
	ChatID     uint
	SenderID   uint
	ReceiverID uint
	isRead     bool
}
