package models

type User struct {
	// `` : go 에서 struct 묶을 때 백틱 사용
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `gorm:"unique" json:"email"`
}
