package gosql

import (
	"fmt"
	. "github.com/suboat/go-sql-kit"
	. "github.com/suboat/go-sql-kit/sql"
)

type SQLXQuery struct {
	SQLQuery
	index  int
	values []interface{}
}

func NewSQLXQuery(size ...int) *SQLXQuery {
	cnt := 10
	if len(size) != 0 {
		cnt = size[0]
	}
	s := &SQLXQuery{SQLQuery: *NewSQLQuery(), index: 1, values: make([]interface{}, 0, cnt)}
	s.SetValueFormat(s.ValueFormat)
	return s
}

func (s *SQLXQuery) ValueFormat(key string, field string, value interface{}) string {
	opera := ""
	switch key {
	case QueryKeyEq:
		opera = "="
	case QueryKeyNe:
		opera = "<>"
	case QueryKeyLt:
		opera = "<"
	case QueryKeyLte:
		opera = "<="
	case QueryKeyGt:
		opera = ">"
	case QueryKeyGte:
		opera = ">="
	case QueryKeyLike:
		return fmt.Sprintf("%v LIKE '%%%v%%'", field, value)
	}
	defer func() {
		s.index++
		s.values = append(s.values, value)
	}()
	return fmt.Sprintf("%v%v$%v", field, opera, s.index)
}

func (s *SQLXQuery) JSONtoSQLString(str string) (string, []interface{}, error) {
	if err := s.ParseJSONString(str); err != nil {
		return "", nil, err
	}
	return s.String(), s.values, nil
}

func (s *SQLXQuery) SQLString(m map[string]interface{}) (string, []interface{}, error) {
	if err := s.Parse(m); err != nil {
		return "", nil, err
	}
	return s.String(), s.values, nil
}
