package repository

import "errors"

const (
	errMsgChatNotFound = "chat not found"
	errMsgUserNotFound = "user not found"
)

var (
	ErrUserNotFound = errors.New(errMsgUserNotFound) // ErrUserNotFound signal error due user not found.
	ErrChatNotFound = errors.New(errMsgChatNotFound) // ErrChatNotFound signal error due chat not found.
)
