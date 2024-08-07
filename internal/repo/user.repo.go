package repo

type UserRepo struct{}

func NewNewRepo() *UserRepo {
	return &UserRepo{}
}

// user repo u
func (ur *UserRepo) GetInfoUser() string {
	return "Tips"
}
