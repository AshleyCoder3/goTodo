package model

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
	Body   string `json:"body"`
	UserID string `json:"userId"`
}
type NewTodo struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID string `json:"userId"`
}
