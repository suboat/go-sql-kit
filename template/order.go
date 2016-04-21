package gosql

import (
	. "github.com/suboat/go-sql-kit"
	"strings"
)

type TemplateOrder struct {
	RuleMapping
	OrderRoot
}

func NewTemplateOrder() *TemplateOrder {
	return new(TemplateOrder)
}

func (s *TemplateOrder) String() string {
	return ""
}

func (s *TemplateOrder) TemplateString(v interface{}) (string, error) {
	return s.String(), nil
}
