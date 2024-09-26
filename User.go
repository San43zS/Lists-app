package Lists_app

type User struct {
	Id       int    `json:"-"` // we don't need this field in JSON
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
