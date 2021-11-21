package todo

// Gin binding in Go tutorial.
// https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
