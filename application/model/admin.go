package model

type Admin struct {
	ID         uint64
	Username   string
	Password   string
	Permission string
	Email      string
	Phone      string
	Call       string
	Avatar     string
	Status     bool
	CreateTime uint64 `gorm:"autoCreateTime"`
	UpdateTime uint64 `gorm:"autoUpdateTime"`
}
