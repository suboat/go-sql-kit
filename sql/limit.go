package gosql

import (
	"fmt"
	. "github.com/suboat/go-sql-kit"
)

type SQLLimit struct {
	RuleLimit
	LimitRoot
	LimitValueSet
}

func NewSQLLimit() *SQLLimit {
	return new(SQLLimit)
}

func (s *SQLLimit) String() string {
	if s.Values == nil || len(s.Values) == 0 {
		return ""
	}
	for _, iv := range s.Values {
		if v, ok := iv.(*LimitValue); ok && v.IsLimited() {
			switch v.Key {
			case LimitKeyLimit:
				s.Limit = s.GetLimit(v.Value)
			case LimitKeySkip:
				s.Skip = v.Value
			case LimitKeyPage:
				s.Page = v.Value
			}
		}
	}
	if s.LimitValueSet.IsLimited() {
		return s.ValueString()
	}
	return ""
}

func (s *SQLLimit) ValueString() string {
	if s.Limit <= 0 {
		return ""
	}
	sql := fmt.Sprintf(`LIMIT %v`, s.Limit)
	if skip := s.Skip + s.Limit*s.Page; skip > 0 {
		sql += fmt.Sprintf(` OFFSET %v`, skip)
	}
	return sql
}

func (s *SQLLimit) JSONtoSQLString(str string) (string, error) {
	if err := s.ParseJSONString(str); err != nil {
		return "", err
	}
	return s.String(), nil
}

func (s *SQLLimit) SQLString(m map[string]interface{}) (string, error) {
	if err := s.Parse(m); err != nil {
		return "", err
	}
	return s.String(), nil
}
