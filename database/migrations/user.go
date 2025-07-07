package migrations

type User struct {
	ID       uint `gorm:"primaryKey"`
	Fullname string
	Email    string
	Password string
	Status   string
}
