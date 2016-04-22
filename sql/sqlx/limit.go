package gosql

import . "github.com/suboat/go-sql-kit/sql"

type SQLXLimit struct {
	SQLLimit
}

func NewSQLXLimit() *SQLXLimit {
	return &SQLXLimit{SQLLimit: *NewSQLLimit()}
}
