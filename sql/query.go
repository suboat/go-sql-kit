package gosql

import . "github.com/suboat/go-sql-kit"

type SQLQueryRoot struct {
	QueryRoot
}

func (q *SQLQueryRoot) ParseJSON(str string) error {
	return nil
}

func (q *SQLQueryRoot) QueryKey() string {
	return QueryKey1_and
}

func (q *SQLQueryRoot) QueryValue() []IQuery {
	return []IQuery{q.Value}
}

func (q *SQLQueryRoot) String() string {
	return q.Value.String()
}
