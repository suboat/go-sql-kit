package gosql

import "errors"

var (
	ErrTypeMap    error = errors.New("gosql: interface{} is`t type of map[string]interface{}")
	ErrTypeString error = errors.New("gosql: interface{} is`t type of string")
)
