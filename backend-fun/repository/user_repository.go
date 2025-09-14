package repository

type User struct {
	Username string
	Password string
}

type UserRepository struct {
	users map[string]User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]User),
	}
}

func (r *UserRepository) Save(user User) {
	r.users[user.Username] = user
}

func (r *UserRepository) FindByUsername(username string) (User, bool) {
	user, found := r.users[username]
	return user, found
}
