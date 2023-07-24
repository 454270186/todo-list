package dal

type Users struct {
	ID       int    `gorm:"primary_key"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

type Tasks struct {
	ID        string    `gorm:"primary_key"`
	UserID    int    `gorm:"column:user_id"`
	Name      string `gorm:"column:name"`
	Completed bool   `gorm:"column:completed"`
}
