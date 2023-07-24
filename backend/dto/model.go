package dto

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"msg"`
	UserID     int    `json:"user_id"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRes struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"msg"`
	UserID     int    `json:"user_id"`
}

type Task struct {
	ID        string    `json:"id"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
