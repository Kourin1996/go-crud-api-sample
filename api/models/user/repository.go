package user

type IUserRepository interface {
	Create(*User) (*User, error)
	Get(int64) (*User, error)
	GetByUsername(string) (*User, error)
}
