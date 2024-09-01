package model

import (
	"time"
)

const TableNameUser = "user"

type User struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:primary key" json:"id"`                      // primary key
	Name      string    `gorm:"column:name;not null;default:new user;comment:user name" json:"name"`                        // user name
	Age       int32     `gorm:"column:age;not null;comment:user age" json:"age"`                                            // user age
	Gender    string    `gorm:"column:gender;not null;default:unknown;comment:user gender" json:"gender"`                   // user gender
	Password  string    `gorm:"column:password;not null;comment:user gender" json:"password"`                               // user gender
	NickName  string    `gorm:"column:nick_name;not null;comment:user nick name" json:"nick_name"`                          // user nick name
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:create time" json:"created_at"` // create time
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:update time" json:"updated_at"` // update time
}

func (*User) TableName() string {
	return TableNameUser
}
