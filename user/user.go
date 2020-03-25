package user

// User is test struct.
type User struct {
	Name string
}

// UserRepository is test interface.
type UserRepository interface {
	FindOne(id int) (*User, error)
}
