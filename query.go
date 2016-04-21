package gosql

const (
	QueryKey1_and string = "%and"
	QueryKey1_or         = "%or"
	QueryKey1_in         = "%in"
)

func IsQueryKey1(str string) bool {
	switch str {
	case QueryKey1_and, QueryKey1_or, QueryKey1_in:
		return true
	}
	return false
}

const (
	QueryKey2_eq   string = "%eq"   // 等于
	QueryKey2_ne          = "%ne"   // 不等于
	QueryKey2_lt          = "%lt"   // 小于
	QueryKey2_lte         = "%lte"  // 小于等于
	QueryKey2_gt          = "%gt"   // 大于
	QueryKey2_gte         = "%gte"  // 大于等于
	QueryKey2_like        = "%like" // 模糊搜索
)

func IsQueryKey2(str string) bool {
	switch str {
	case QueryKey2_eq, QueryKey2_ne:
		return true
	case QueryKey2_lt, QueryKey2_lte:
		return true
	case QueryKey2_gt, QueryKey2_gte:
		return true
	case QueryKey2_like:
		return true
	}
	return false
}

type IQuery interface {
	IsAnonymous() bool
	QueryKey() string
	QueryValue() []IQuery
}

type QueryRoot struct {
	Value []IQuery
}

func (q *QueryRoot) IsAnonymous() bool {
	return false
}

func (q *QueryRoot) QueryKey() string {
	return QueryKey1_and
}

func (q *QueryRoot) QueryValue() []IQuery {
	return q.Value
}

type QueryElem struct {
	anonymous bool
	Key       string
	Value     []IQuery
}

func (q *QueryElem) IsAnonymous() bool {
	return q.anonymous
}

func (q *QueryElem) QueryKey() string {
	return q.Key
}

func (q *QueryElem) QueryValue() []IQuery {
	return q.Value
}

type QueryValue struct {
	Key   string
	Field string
	Value string
}

func (q *QueryValue) IsAnonymous() bool {
	return false
}

func (q *QueryValue) QueryKey() string {
	return q.Key
}

func (q *QueryValue) QueryValue() []IQuery {
	return nil
}
