package gosql

import (
	"fmt"
	. "github.com/suboat/go-sql-kit"
	"strings"
)

type SQLQuery struct {
	SQLRule
	QueryRoot
}

func NewSQLQuery() *SQLQuery {
	s := new(SQLQuery).AllowCommon()
	s.SetValueFormat(s.valueFormat)
	return s
}

func (s *SQLQuery) AllowCommon() *SQLQuery {
	s.Allow(QueryKeyAnd, QueryKeyOr,
		QueryKeyEq, QueryKeyNe,
		QueryKeyLt, QueryKeyLte,
		QueryKeyGt, QueryKeyGte,
	)
	return s
}

func (s *SQLQuery) String() string {
	if s.Value == nil || len(s.Value) == 0 {
		return ""
	}
	set := make([]string, 0, len(s.Value))
	for _, iv := range s.Value {
		if v, ok := iv.(*QueryElem); ok {
			if str := s.elemString(v); len(str) != 0 {
				set = append(set, str)
			}
		}
	}
	if len(set) != 0 {
		return "WHERE " + strings.Join(set, " AND ")
	}
	return ""
}

func (s *SQLQuery) elemString(elem *QueryElem) string {
	if !s.IsAllowed(elem.Key) {
		return ""
	}
	set := make([]string, 0, len(elem.Value))
	for _, iv := range elem.Value {
		if v, ok := iv.(*QueryElem); ok {
			if str := s.elemString(v); len(str) != 0 {
				set = append(set, str)
			}
		} else if v, ok := iv.(*QueryValue); ok {
			if str := s.valueString(v); len(str) != 0 {
				set = append(set, str)
			}
		}
	}
	if len(set) == 0 {
		return ""
	} else if elem.IsAnonymous() {
		return strings.Join(set, " AND ")
	} else {
		switch elem.Key {
		case QueryKeyAnd:
			return strings.Join(set, " AND ")
		case QueryKeyOr:
			if len(set) == 1 {
				return set[0]
			}
			return fmt.Sprintf("(%v)", strings.Join(set, " OR "))
		}
	}
	return ""
}

func (s *SQLQuery) valueFormat(key string, field string, value interface{}) string {
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
	switch value.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		return fmt.Sprintf("%v%v%v", field, opera, value)
	default:
		return fmt.Sprintf("%v%v'%v'", field, opera, value)
	}
}

func (s *SQLQuery) valueString(v *QueryValue) string {
	return s.ValueString(v)
}

func (s *SQLQuery) JSONtoSQLString(str string) (string, error) {
	if err := s.ParseJSONString(str); err != nil {
		return "", err
	}
	return s.String(), nil
}

func (s *SQLQuery) SQLString(m map[string]interface{}) (string, error) {
	if err := s.Parse(m); err != nil {
		return "", err
	}
	return s.String(), nil
}
