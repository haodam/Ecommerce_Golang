package accountmodel

import (
	"time"
)

type Account struct {
	ID          int       `json:"id" gorm:"primaryKey autoIncrement:true"`
	AccountId   string    `json:"account_id" gorm:"unique"`
	Username    string    `json:"username" gorm:"column:username"`
	Password    string    `json:"password" gorm:"column:password"`
	Name        string    `json:"name" gorm:"column:name"`
	Email       string    `json:"email" gorm:"column:email"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number"`
	Address     string    `json:"address" gorm:"column:address"`
	Type        int       `json:"type" gorm:"column:type"`
	Status      int       `json:"status" gorm:"column:status"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time"`
}

func (Account) TableName() string { return "accounts" }
