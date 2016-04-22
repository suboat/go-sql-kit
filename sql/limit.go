package gosql

import (
	"fmt"
	. "github.com/suboat/go-sql-kit"
)

type SQLLimit struct {
	RuleLimit
	LimitRoot
	limit int
	skip  int
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
		s.limit = v.Limit
		s.skip = v.Skip + v.Limit*v.Page
		return s.ValueString()
	}
	return ""
}

func (s *SQLLimit) ValueString() string {
	if s.limit <= 0 {
		return ""
	}
	sql := fmt.Sprintf(`LIMIT %v`, s.limit)
	if s.skip > 0 {
		sql += fmt.Sprintf(` OFFSET %v`, s.skip)
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

func (s *SQLLimit) GetLimit() int {
	return s.limit
}

func (s *SQLLimit) GetSkip() int {
	return s.skip
}
