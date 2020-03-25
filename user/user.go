package user

// User
type User struct {
	Name string
}

// UserRepository
type UserRepository interface {
	FindOne(id int) (*User, error)
}