package model

type Users struct {
	Id       uint   `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}
