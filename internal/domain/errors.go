package domain

import "errors"

var (
	ErrInvalidUser      = errors.New("invalid user")
	ErrCannotFriendSelf = errors.New("cannot add yourself as a friend")
	ErrAlreadyFriends   = errors.New("users are already friends")
	ErrNotFriends       = errors.New("users are not friends")
)
