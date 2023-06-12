package src

import "gorm.io/gorm"

type Envs struct {
	DbUrl     string
	Token     string
	AppId     int32
	AppHash   string
	Encyptkey string
}
type Accounts struct {
	gorm.Model
	UserID        int64 `gorm:"primaryKey"`
	OwnerId       int64
	StringSession string
	AppId         int32
	AppHash       string
}
