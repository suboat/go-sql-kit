package gosql

import "errors"

var (
	ErrTypeNil    error = errors.New("gosql: interface{} must not be nil")
	ErrTypeMap          = errors.New("gosql: interface{} is`t type of map[string]interface{}")
	ErrTypeString       = errors.New("gosql: interface{} is`t type of string")
)
