package dal

type Users struct {
	ID       int    `gorm:"primary_key"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

type Tasks struct {
	ID          int    `gorm:"primary_key"`
	UserID      int    `gorm:"column:user_id"`
	Title       string `gorm:"column:title"`
	IsCompleted bool   `gorm:"column:is_completed"`
}
