package gosql

import . "github.com/suboat/go-sql-kit"

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
