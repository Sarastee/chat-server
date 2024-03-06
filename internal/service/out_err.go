package service

import "errors"

const errMsgUserNotInTheChat = "user not in chat"

var ErrMsgUserNotInTheChat = errors.New(errMsgUserNotInTheChat) // ErrMsgUserNotInTheChat signal error due user not in chat.
