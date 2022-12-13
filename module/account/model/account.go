package accountmodel

import (
	"errors"
	"strings"
	"time"
)

type Account struct {
	ID          int       `json:"id" gorm:"column:id, primaryKey, autoIncrement:true"`
	AccountId   string    `json:"account_id" gorm:"unique"`
	Username    string    `json:"username" gorm:"column:username"`
	Password    string    `json:"password" gorm:"column:password"`
	Name        string    `json:"name" gorm:"column:name"`
	Email       string    `json:"email" gorm:"column:email"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number"`
	Address     string    `json:"address" gorm:"column:address"`
	Type        int       `json:"type" gorm:"column:type"`
	Status      string    `json:"status" gorm:"column:status"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time"`
}

func (Account) TableName() string { return "accounts" }

func (data *Account) Validate() error {
	data.AccountId = strings.TrimSpace(data.AccountId)
	data.Name = strings.TrimSpace(data.Name)

	if data.AccountId == "" {
		return errors.New("account ID can not be empty")
	}
	if data.Name == "" {
		return errors.New("name can not be empty")
	}
	return nil
}
