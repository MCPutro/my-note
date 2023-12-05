package entity

type User struct {
	ID       string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Token    string `gorm:"-"`
}
