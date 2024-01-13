package models

type User struct {
	Id       string `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"not null"`
	Password string `json:"password" gorm:"type:varchar(16);not null"`
	Money    int    `json:"money" gorm:"size:32;default:0"`
	Ranking  int    `json:"ranking" gorm:"size:32";default:1000`
}
