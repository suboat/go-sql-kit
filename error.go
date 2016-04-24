package gosql

import "errors"

var (
	ErrTypeNil    error = errors.New("gosql: interface{} must not be nil")
	ErrTypeMap          = errors.New("gosql: interface{} must be type of map[string]interface{}")
	ErrTypeString       = errors.New("gosql: interface{} must be type of string")
	ErrTypeInt          = errors.New("gosql: interface{} must be type of int")
	ErrTypeValue        = errors.New("gosql: invalid value")
)
