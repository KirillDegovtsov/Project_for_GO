package domain

import "errors"

var (
	NoteNotFound      = errors.New("note not found")
	NoteAlreadyExists = errors.New("task already exists")
	UserNotFound      = errors.New("user not found")
	UserAlreadyExists = errors.New("user already exists")
	InvalidPassword   = errors.New("invalid password")
	BadRequest        = errors.New("bad request")
	InternalError     = errors.New("internal error")
	Unauthorized      = errors.New("unauthorized")
	KeyAlreadyExists  = errors.New("key already exists")
	InvalidData       = errors.New("invalid data")
)
