package model

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(100);not null"`
	Phone    string `gorm:"type:varchar(20);not null"`
}
