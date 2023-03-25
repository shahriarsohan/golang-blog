package models

type Post struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Image       string `json:"image"`
	AuthorID    string `json:"author"`
	User        User   `gorm:"foreignKey:AuthorID"`
}
