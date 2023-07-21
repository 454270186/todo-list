package dto

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"msg"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRes struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"msg"`
}
