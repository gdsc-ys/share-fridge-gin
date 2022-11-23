package model

type User struct {
	ID      string    `json:"id" gorm:"primaryKey"`
	Type    string    `json:"type" gorm:"primaryKey"`
	Name    string    `json:"nickname"`
	Image   string    `json:"profile_image"`
	Fridges []*Fridge `gorm:"many2many:favorite;"`
}
