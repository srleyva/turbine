package user

import (
	"errors"
	proto "github.com/srleyva/turbine/user-service/proto/user"
)

// Primarily for testing as the storage is transient
type DataStore struct {
	users []proto.User
}

func (d *DataStore) Create(user *proto.User) (*proto.User, error) {
	update := append(d.users, *user)
	d.users = update
	return user, nil
}

func (d *DataStore) GetAll() (*[]proto.User, error) {
	var users []proto.User = d.users
	return &users, nil
}

func (d *DataStore) Get(username string) (*proto.User, error) {
	// as this is only for testing not worried about introducing a efficent algorithm
	for _, user := range d.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}
