package user

import "time"

// UserDTO  表示输出到外部的用户信息
type UserDTO struct {
	ID int64
	UserName string
	CreatedAt time.Time
}
