package model

import (
	"encoding/json"
	"time"

	"github.com/andresmeireles/resume/utils"
)

type User struct {
	Id        *uint      `db:"id"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Password  string     `db:"password" json:"-"`
	Hash      *string    `db:"hash"`
	ExpiresAt *time.Time `db:"expires_at"`
}

func (u User) GetId() *uint {
	return u.Id
}

func (u User) TableName() string {
	return "users"
}

func (u User) InsertFields() map[string]interface{} {
	return map[string]interface{}{
		"name":       u.Name,
		"email":      u.Email,
		"password":   u.Password,
		"expires_at": u.ExpiresAt,
		"hash":       u.Hash,
	}
}

func (u *User) SetHash() (*User, error) {
	expiresAt := time.Now().Add(time.Hour * 24 * 30)
	hash, hashErr := utils.PasswordHash(u.Name + u.Password + expiresAt.String())

	if hashErr != nil {
		return nil, hashErr
	}

	u.Hash = &hash
	u.ExpiresAt = &expiresAt
	return u, nil
}

func (u *User) IsHashExpired() bool {
	return time.Now().After(*u.ExpiresAt)
}

func (u *User) GetData() string {
	data := map[string]string{
		"name":  u.Name,
		"email": u.Email,
	}
	json, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(json)
}

// This function must not be used as return json, it exposes secret information
func (u User) Content() map[string]interface{} {
	return map[string]interface{}{
		"name":       u.Name,
		"email":      u.Email,
		"password":   u.Password,
		"hash":       u.Hash,
		"expires_at": u.ExpiresAt,
	}
}
