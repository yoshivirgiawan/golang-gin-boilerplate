package user

import "time"

type User struct {
	ID             string
	Name           string
	Email          string
	Password       string
	AvatarFileName string
	Role           string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
