package Lists_app

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
