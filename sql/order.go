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

func (s *SQLOrder) String(alias ...string) string {
	if s.Values == nil || len(s.Values) == 0 {
		return ""
	}
	set := make([]string, 0, len(s.Values))
	for _, iv := range s.Values {
		if v, ok := iv.(*OrderValue); ok {
			if str := s.valueString(v, alias...); len(str) != 0 {
				set = append(set, str)
			}
		}
	}
	if len(set) == 0 {
		return ""
	}
	return "ORDER BY " + strings.Join(set, ", ")
}

func (s *SQLOrder) valueString(v *OrderValue, alias ...string) string {
	if v == nil {
	} else if !s.IsAllowed(v.Field) {
	} else if v.Field = s.GetMapping(v.Field); len(v.Field) != 0 {
		if f, ok := s.GetRuleMappingResult(v.Field); ok {
			if result, ok := f(v.Field, v.IsASC(), v.Key, alias...); ok {
				if str, ok := result.(string); ok {
					return str
				}
			}
		}
		if len(alias) != 0 {
			v.Field = fmt.Sprintf("%v.%v", alias[0], v.Field)
		}
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

func (s *SQLOrder) JSONtoSQLString(str string, alias ...string) (string, error) {
	if err := s.ParseJSONString(str); err != nil {
		return "", err
	}
	return s.String(alias...), nil
}

func (s *SQLOrder) SQLString(m map[string]interface{}, alias ...string) (string, error) {
	if err := s.Parse(m); err != nil {
		return "", err
	}
	return s.String(alias...), nil
}
