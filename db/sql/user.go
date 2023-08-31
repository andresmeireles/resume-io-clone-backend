package sql

type User struct {
	ID        uint `gorm:"primaryKey"`
	name      string
	email     string
	password  string
	hash      string
	expiresAt int
}

func (u User) create(name, email, password string) (*User, error) {
	user := User{
		name:      name,
		email:     email,
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

func (u User) getById(id int) (*User, error) {
	db := Db()
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u User) update(id int, name, email string) (*User, error) {
	db := Db()
	var user User
	result := db.Find(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if name != "" && user.name != name {
		user.name = name
	}
	if email != "" && user.email != email {
		user.email = email
	}
	db.Save(&user)
	return &user, nil
}
