package gosql

import (
	"fmt"
	. "github.com/suboat/go-sql-kit"
)

type SQLLimit struct {
	RuleLimit
	LimitRoot
}

func NewSQLLimit() *SQLLimit {
	return new(SQLLimit)
}

func (s *SQLLimit) String() string {
	if s.Value == nil {
	} else if v, ok := s.Value.(*LimitValue); ok {
		return s.valueString(v)
	}
	return ""
}

func (s *SQLLimit) valueString(v *LimitValue) string {
	if v == nil {
	} else if v.Limit = s.Limit(v.Limit); v.IsLimited() {
		return s.ValueString(v)
	}
	return ""
}

func (s *SQLLimit) ValueString(v *LimitValue) string {
	sql := fmt.Sprintf(`LIMIT %v`, v.Limit)
	v.Skip = v.Skip + v.Limit*v.Page
	if v.Skip > 0 {
		sql += fmt.Sprintf(` OFFSET %v`, v.Skip)
	}
	return sql
}

func (s *SQLLimit) JSONtoSQLString(str string) (string, error) {
	if err := s.ParseJSONString(str); err != nil {
		return "", err
	}
	return s.String(), nil
}

func (s *SQLLimit) SQLString(str string) (string, error) {
	if err := s.Parse(str); err != nil {
		return "", err
	}
	return s.String(), nil
}
