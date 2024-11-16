package user

type User struct {
	Id       int    `json:"-"` // we don't need this field in JSON
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SingUpReq struct{}

type SingInReq struct{}
