package gosql

import . "github.com/suboat/go-sql-kit"

type SQLXParser struct {
	Parser
	query *SQLXQuery
	order *SQLXOrder
	limit *SQLXLimit
}

func NewSQLXParser() *SQLXParser {
	return new(SQLXParser)
}

func (s *SQLXParser) InitALL() *SQLXParser {
	return s.SetQuery(NewSQLXQuery()).SetOrder(NewSQLXOrder()).SetLimit(NewSQLXLimit())
}

func (s *SQLXParser) SetQuery(obj *SQLXQuery) *SQLXParser {
	s.query = obj
	s.Add(obj)
	return s
}

func (s *SQLXParser) GetQuery() *SQLXQuery {
	return s.query
}

func (s *SQLXParser) SetOrder(obj *SQLXOrder) *SQLXParser {
	s.order = obj
	s.Add(obj)
	return s
}

func (s *SQLXParser) GetOrder() *SQLXOrder {
	return s.order
}

func (s *SQLXParser) SetLimit(obj *SQLXLimit) *SQLXParser {
	s.limit = obj
	s.Add(obj)
	return s
}

func (s *SQLXParser) GetLimit() *SQLXLimit {
	return s.limit
}
