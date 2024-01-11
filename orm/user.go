package orm

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);not null"`
	Nickname  string `gorm:"type:varchar(20);not null"`
	Password  string `gorm:"type:varchar(100);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
}

func (User) TableName() string {
	return "user"
}
