package domain

import "time"

type User struct {
	id          string
	username    string
	email       string
	displayName string
	bio         string
	profilePic  string
	dateJoined  time.Time
	isPrivate   bool

	friends map[string]*User
}

func NewUser(
	id string,
	username string,
	email string,
	displayName string,
	isPrivate bool,
) *User {

	return &User{
		id:          id,
		username:    username,
		email:       email,
		displayName: displayName,
		isPrivate:   isPrivate,
		dateJoined:  time.Now(),
		friends:     make(map[string]*User),
	}
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) DisplayName() string {
	return u.displayName
}

func (u *User) AddFriend(other *User) error {
	if other == nil {
		return ErrInvalidUser
	}

	if other.id == u.id {
		return ErrCannotFriendSelf
	}

	if u.IsFriendsWith(other) {
		return nil
	}

	u.friends[other.id] = other
	other.friends[u.id] = u

	return nil
}

func (u *User) RemoveFriend(other *User) {
	if other == nil {
		return
	}

	delete(u.friends, other.id)
	delete(other.friends, u.id)
}

func (u *User) IsFriendsWith(other *User) bool {
	if other == nil {
		return false
	}

	_, ok := u.friends[other.id]
	return ok
}

func (u *User) CanView(viewer *User) bool {
	// Self access
	if viewer != nil && viewer.id == u.id {
		return true
	}

	// Public profile
	if !u.isPrivate {
		return true
	}

	// Private but friend
	if viewer != nil && u.IsFriendsWith(viewer) {
		return true
	}

	return false
}
