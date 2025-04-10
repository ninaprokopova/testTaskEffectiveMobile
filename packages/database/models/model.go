package database

type User struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"size:100;not null"`
	Surname    string `gorm:"size:100;not null"`
	Patronymic string `gorm:"size:100"`
	Age        int
	Gender     string `gorm:"size:100"`
	Nation     string `gorm:"size:100"`
}
