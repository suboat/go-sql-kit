package gosql

import (
	"bytes"
	"fmt"
	. "github.com/suboat/go-sql-kit"
	. "github.com/suboat/go-sql-kit/sql"
)

type SQLXQuery struct {
	SQLQuery
	index  int
	values []interface{}
}

func NewSQLXQuery() *SQLXQuery {
	s := &SQLXQuery{SQLQuery: *NewSQLQuery(), index: 1, values: nil}
	s.SetValueFormat(s.ValueString)
	return s
}

func (s *SQLXQuery) String(alias ...string) string {
	s.index = 1
	s.values = nil
	return s.SQLQuery.String(alias...)
}

func (s *SQLXQuery) ValueString(v *QueryValue) string {
	opera := ""
	switch v.Key {
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
		return fmt.Sprintf("%v LIKE '%%%v%%'", v.Field, v.Value)
	case QueryKeyIn:
		if vs, ok := v.Value.([]interface{}); ok {
			if vs == nil || len(vs) == 0 {
				return ""
			}
			var sb bytes.Buffer
			sb.WriteString(fmt.Sprintf("%v IN (", v.Field))
			l := len(vs)
			for i, vi := range vs {
				sb.WriteString(fmt.Sprintf("$%v", s.index))
				s.index++
				s.values = append(s.values, vi)
				if i+1 < l {
					sb.WriteString(", ")
				}
			}
			sb.WriteString(")")
			return sb.String()
		}
		return ""
	case QueryKeyBetween, QueryKeyNotBetween:
		if vs, ok := v.Value.([]interface{}); ok {
			if vs == nil || len(vs) < 2 {
				return ""
			}
			defer func() {
				s.index++
				s.values = append(s.values, vs[0])
				s.index++
				s.values = append(s.values, vs[1])
			}()
			if v.Key == QueryKeyBetween {
				return fmt.Sprintf("%v BETWEEN $%v AND $%v", v.Field, s.index, s.index+1)
			} else {
				return fmt.Sprintf("%v NOT BETWEEN $%v AND $%v", v.Field, s.index, s.index+1)
			}
		}
	default:
		return ""
	}
	defer func() {
		s.index++
		s.values = append(s.values, v.Value)
	}()
	return fmt.Sprintf("%v%v$%v", v.Field, opera, s.index)
}

func (s *SQLXQuery) GetValues() []interface{} {
	return s.values
}

func (s *SQLXQuery) JSONtoSQLString(str string, alias ...string) (string, []interface{}, error) {
	if err := s.ParseJSONString(str); err != nil {
		return "", nil, err
	}
	return s.String(alias...), s.GetValues(), nil
}

func (s *SQLXQuery) SQLString(m map[string]interface{}, alias ...string) (string, []interface{}, error) {
	if err := s.Parse(m); err != nil {
		return "", nil, err
	}
	return s.String(alias...), s.GetValues(), nil
}
