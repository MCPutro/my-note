package entity

import "time"

type Note struct {
	ID          int
	Text        string
	Visible     bool
	UserId      string
	User        User
	ThemeColor  string
	CreatedDate time.Time
	UpdateDate  time.Time
}
