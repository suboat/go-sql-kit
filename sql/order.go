package gosql

import (
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
	if s.Value == nil || len(s.Value) == 0 {
		return ""
	}
	set := make([]string, 0, len(s.Value))
	for _, iv := range s.Value {
		if v, ok := iv.(*OrderValue); ok {
			if s.IsAllowed(v.Field) {
				if v.IsASC() {
					set = append(set, s.GetMapping(v.Field)+" ASC")
				} else {
					set = append(set, s.GetMapping(v.Field)+" DESC")
				}
			}
		}
	}
	if len(set) == 0 {
		return ""
	}
	return "ORDER BY " + strings.Join(set, ", ")
}

func (s *SQLOrder) JSONStringToSQL(str string) (string, error) {
	if err := s.ParseJSONString(str); err != nil {
		return "", err
	}
	return s.String(), nil
}

func (s *SQLOrder) JSONStringsToSQL(strs []string) (string, error) {
	if err := s.ParseJSONStrings(strs); err != nil {
		return "", err
	}
	return s.String(), nil
}
