package gosql

import (
	"fmt"
	. "github.com/suboat/go-sql-kit"
	"strings"
)

type TemplateQuery struct {
	RuleMapping
	QueryRoot
}

func NewTemplateQuery() *TemplateQuery {
	return new(TemplateQuery)
}

func (s *TemplateQuery) String() string {
	return ""
}

func (s *TemplateQuery) TemplateString(v interface{}) (string, error) {
	return s.String(), nil
}
