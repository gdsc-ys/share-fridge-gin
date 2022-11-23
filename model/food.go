package model

type Food struct {
	ID             uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string `json:"name"`
	Category       string `json:"category"`
	ExpirationDate string `json:"expiration_date"`
	Image          string `json:"image"`
	Sharer         string `json:"sharer"`
}
