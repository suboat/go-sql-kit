package gosql

import (
	"fmt"
	. "github.com/suboat/go-sql-kit"
	"strings"
)

type ValueStringFunc func(*QueryValue) string

type SQLQuery struct {
	RuleMapping
	QueryRoot
	valueStringFunc ValueStringFunc
}

func NewSQLQuery() *SQLQuery {
	return new(SQLQuery).AllowCommonKey().SetValueFormat(nil)
}

func (s *SQLQuery) AllowCommonKey() *SQLQuery {
	s.Allow(QueryKeyAnd, QueryKeyOr,
		QueryKeyEq, QueryKeyNe,
		QueryKeyLt, QueryKeyLte,
		QueryKeyGt, QueryKeyGte,
	)
	return s
}

func (s *SQLQuery) SetValueFormat(f ValueStringFunc) *SQLQuery {
	if f != nil {
		s.valueStringFunc = f
	} else {
		s.valueStringFunc = s.ValueString
	}
	return s
}

func (s *SQLQuery) String() string {
	if s.Values == nil || len(s.Values) == 0 {
		return ""
	}
	set := make([]string, 0, len(s.Values))
	for _, iv := range s.Values {
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
	set := make([]string, 0, len(elem.Values))
	for _, iv := range elem.Values {
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

func (s *SQLQuery) valueString(v *QueryValue) string {
	if v == nil {
	} else if !s.IsAllowed(v.Field) {
	} else {
		v.Field = s.GetMapping(v.Field)
		return s.valueStringFunc(v)
	}
	return ""

}

func (s *SQLQuery) ValueString(v *QueryValue) string {
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
	default:
		return ""
	}
	switch v.Value.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		return fmt.Sprintf("%v%v%v", v.Field, opera, v.Value)
	default:
		return fmt.Sprintf("%v%v'%v'", v.Field, opera, v.Value)
	}
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
