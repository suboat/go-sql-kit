package gosql

import (
	. "github.com/suboat/go-sql-kit"
	"strings"
)

type SQLParser struct {
	Parser
	query *SQLQuery
	order *SQLOrder
	limit *SQLLimit
}

func NewSQLParser() *SQLParser {
	return new(SQLParser)
}

func (s *SQLParser) InitALL() *SQLParser {
	return s.SetQuery(NewSQLQuery()).SetOrder(NewSQLOrder()).SetLimit(NewSQLLimit())
}

func (s *SQLParser) SetQuery(obj *SQLQuery) *SQLParser {
	s.query = obj
	s.Add(obj)
	return s
}

func (s *SQLParser) GetQuery() *SQLQuery {
	return s.query
}

func (s *SQLParser) SetOrder(obj *SQLOrder) *SQLParser {
	s.order = obj
	s.Add(obj)
	return s
}

func (s *SQLParser) GetOrder() *SQLOrder {
	return s.order
}

func (s *SQLParser) SetLimit(obj *SQLLimit) *SQLParser {
	s.limit = obj
	s.Add(obj)
	return s
}

func (s *SQLParser) GetLimit() *SQLLimit {
	return s.limit
}

func (s *SQLParser) JoinString(query bool, order bool, limit bool) string {
	set := make([]string, 0, 3)
	if !query {
	} else if str := s.GetQuery().String(); len(str) == 0 {
		return ""
	} else {
		set = append(set, str)
	}
	if !order {
	} else if str := s.GetOrder().String(); len(str) == 0 {
		return ""
	} else {
		set = append(set, str)
	}
	if !limit {
	} else if str := s.GetLimit().String(); len(str) == 0 {
		return ""
	} else {
		set = append(set, str)
	}
	return strings.Join(set, " ")
}
