package gosql

import . "github.com/suboat/go-sql-kit/sql"

type SQLXOrder struct {
	SQLOrder
}

func NewSQLXOrder() *SQLXOrder {
	return &SQLXOrder{SQLOrder: *NewSQLOrder()}
}
