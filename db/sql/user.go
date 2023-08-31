package sql

import "github.com/andresmeireles/resume/utils"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	password  string
	hash      string
	expiresAt int
}

func (u User) create(name, email, password string) (*User, error) {
	user := User{
		Name:      name,
		Email:     email,
		password:  password,
		hash:      "",
		expiresAt: 0,
	}
	db := Db()
	create := db.Create(&user)
	if create.Error != nil {
		return nil, create.Error
	}
	return &user, nil
}

func (u User) getAll() ([]User, error) {
	db := Db()
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (u User) GetById(id int) (*User, error) {
	db := Db()
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u User) FindByEmail(email string) (*User, error) {
	db := Db()
	var user User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u User) VerifyPassword(password string) bool {
	hashPassword, err := utils.PasswordHash(password)
	if err != nil {
		return false
	}
	return hashPassword == u.hash
}
