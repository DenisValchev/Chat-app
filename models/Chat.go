package models

import "github.com/jinzhu/gorm"

type Chat struct {
	gorm.Model
	Sender     User `gorm:"foreignkey:SenderID"`
	SenderID   uint
	Receiver   User `gorm:"foreignkey:SenderID"`
	ReceiverID uint
}
