package gosql

import (
	"fmt"
	. "github.com/suboat/go-sql-kit"
	"strings"
)

type SQLOrder struct {
	RuleMapping
	OrderRoot
}

func NewSQLOrder() *SQLOrder {
	return new(SQLOrder)
}

func (s *SQLOrder) String() string {
	if s.Values == nil || len(s.Values) == 0 {
		return ""
	}
	set := make([]string, 0, len(s.Values))
	for _, iv := range s.Values {
		if v, ok := iv.(*OrderValue); ok {
			if str := s.valueString(v); len(str) != 0 {
				set = append(set, str)
			}
		}
	}
	if len(set) == 0 {
		return ""
	}
	return "ORDER BY " + strings.Join(set, ", ")
}

func (s *SQLOrder) valueString(v *OrderValue) string {
	if v == nil {
	} else if !s.IsAllowed(v.Field) {
	} else {
		v.Field = s.GetMapping(v.Field)
		return s.ValueString(v)
	}
	return ""
}

func (s *SQLOrder) ValueString(v *OrderValue) string {
	if v.IsASC() {
		return fmt.Sprintf("%v ASC", v.Field)
	}
	return fmt.Sprintf("%v DESC", v.Field)
}

func (s *SQLOrder) JSONtoSQLString(str string) (string, error) {
	if err := s.ParseJSONString(str); err != nil {
		return "", err
	}
	return s.String(), nil
}

func (s *SQLOrder) SQLString(m map[string]interface{}) (string, error) {
	if err := s.Parse(m); err != nil {
		return "", err
	}
	return s.String(), nil
}
