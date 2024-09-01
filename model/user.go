package model

import (
	"time"
)

const TableNameUser = "user"

type User struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement:true;comment:primary key" json:"id"`                      // primary key
	Username  string    `gorm:"column:username;not null;comment:user name" json:"username"`                                 // user name
	Age       int       `gorm:"column:age;not null;comment:user age" json:"age"`                                            // user age
	Gender    string    `gorm:"column:gender;not null;default:unknown;comment:user gender" json:"gender"`                   // user gender
	Password  string    `gorm:"column:password;not null;comment:user gender" json:"password"`                               // user gender
	Nickname  string    `gorm:"column:nickname;not null;comment:user nickname" json:"nickname"`                             // user nick name
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:create time" json:"created_at"` // create time
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:update time" json:"updated_at"` // update time
}

func (*User) TableName() string {
	return TableNameUser
}
