package auth

// import (
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/andresmeireles/resume/db/sql"
// )

// func TestLoginWithHash(t *testing.T) {
// 	// Test case: Valid hash, user expires in the future
// 	hash := "valid_hash"
// 	user := &sql.User{
// 		ID:        1,
// 		ExpiresAt: int(time.Now().Add(time.Hour).Unix()),
// 	}
// 	sql.User{}.GetByHash = func(hash string) (*sql.User, error) {
// 		if hash == "valid_hash" {
// 			return user, nil
// 		}
// 		return nil, errors.New("Invalid hash")
// 	}

// 	result, err := LoginWithHash(hash)
// 	if err != nil {
// 		t.Errorf("LoginWithHash() returned error: %v", err)
// 	}
// 	if result != user {
// 		t.Errorf("LoginWithHash() returned incorrect user, got: %v, want: %v", result, user)
// 	}
// }

// func TestLoginWithHash_InvalidHash(t *testing.T) {
// 	// Test case: Invalid hash
// 	hash := "invalid_hash"
// 	sql.User{}.GetByHash = func(hash string) (*sql.User, error) {
// 		return nil, errors.New("Invalid hash")
// 	}

// 	result, err := LoginWithHash(hash)
// 	if err == nil {
// 		t.Errorf("LoginWithHash() should have returned an error")
// 	}
// 	if result != nil {
// 		t.Errorf("LoginWithHash() should have returned nil user, got: %v", result)
// 	}
// }

// func TestLoginWithHash_UserExpired(t *testing.T) {
// 	// Test case: Valid hash, user expired
// 	hash := "valid_hash"
// 	user := &sql.User{
// 		ID:        1,
// 		ExpiresAt: time.Now().Add(-time.Hour).Unix(),
// 	}
// 	sql.User{}.GetByHash = func(hash string) (*sql.User, error) {
// 		if hash == "valid_hash" {
// 			return user, nil
// 		}
// 		return nil, errors.New("Invalid hash")
// 	}

// 	result, err := LoginWithHash(hash)
// 	if err == nil {
// 		t.Errorf("LoginWithHash() should have returned an error")
// 	}
// 	if result != nil {
// 		t.Errorf("LoginWithHash() should have returned nil user, got: %v", result)
// 	}
// }
