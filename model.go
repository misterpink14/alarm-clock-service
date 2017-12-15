package main

type Model interface {
	Get()
	Save()
	Delete()
}

type User struct {
	Name         string
	Password     string
	DisplayName  string
	Created      string
	LastModified string
}

func (u *User) Get() (User, error) {
	databaseWrapper := GetDatabaseWrapper()
	return User{}, nil
}

func (u *User) Save() (bool, error) {
	databaseWrapper := GetDatabaseWrapper()
	return false, nil
}

func (u *User) Delete() (bool, error) {
	databaseWrapper := GetDatabaseWrapper()
	return false, nil
}
