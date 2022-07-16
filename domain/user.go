package domain

type User struct {
	id string 
	name string
	level int
	password string
}

func NewUser(id string, name string, level int, password string) *User {
	user := User{id, name, level, password}
	return &user
}

func (user *User) Id() string{
	return user.id
}

func (user *User) Name() string{
	return user.name
}

func (user *User) Level() int{
	return user.level
}

func (user *User) Password() string{
	return user.password
}
