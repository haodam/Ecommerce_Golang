package accountmodel

import "time"

type Account struct {
	ID      int       `json:"id" gorm:"column:id"`
	User    string    `json:"user" gorm:"column:user"`
	Name    string    `json:"name" gorm:"column:name"`
	Email   string    `json:"email" gorm:"column:email"`
	Phone   int       `json:"phone" gorm:"column:phone"`
	Address string    `json:"address" gorm:"column:address"`
	Type    string    `json:"type" gorm:"column:type"`
	Status  string    `json:"status" gorm:"column:status"`
	Created time.Time `json:"created" gorm:"column:created"`
}

type AccountTable struct {
	ID      int       `gorm:"primaryKey"`
	User    string    `sql:"type:CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Name    string    `sql:"type:CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Email   string    `sql:"type:CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Phone   string    `sql:"type:CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Address string    `sql:"type:CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Type    string    `sql:"type:CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Status  string    `sql:"type:CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Created time.Time `sql:"type:CHARACTER SET utf8 COLLATE utf8_general_ci"`
}

func (AccountTable) TableName() string { return "account" }
