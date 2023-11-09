package model

import "time"

type User struct {
	ID        uint64    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	Email     string    `xorm:"not null VARCHAR(100) unique index email"`
}

func (u User) TableName() string {
	return "t_users"
}
