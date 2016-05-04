package gosql

const (
	QueryKeyAnd string = "%and" // AND 与
	QueryKeyOr         = "%or"  // OR  或
)

const (
	QueryKeyEq         string = "%eq"   // 等于
	QueryKeyNe                = "%ne"   // 不等于
	QueryKeyLt                = "%lt"   // 小于
	QueryKeyLte               = "%lte"  // 小于等于
	QueryKeyGt                = "%gt"   // 大于
	QueryKeyGte               = "%gte"  // 大于等于
	QueryKeyLike              = "%like" // 模糊搜索
	QueryKeyIn                = "%in"   // TODO: 暂时不支持
	QueryKeyBetween           = "%bt"
	QueryKeyNotBetween        = "%nbt"
)

func IsQueryKey(str string) bool {
	switch str {
	case QueryKeyAnd, QueryKeyOr:
		return true
	}
	return IsQueryAnonymousKey(str)
}

func IsQueryAnonymousKey(str string) bool {
	switch str {
	case QueryKeyEq, QueryKeyNe:
		return true
	case QueryKeyLt, QueryKeyLte:
		return true
	case QueryKeyGt, QueryKeyGte:
		return true
	case QueryKeyLike:
		return true
	case QueryKeyBetween, QueryKeyNotBetween:
		return true
	}
	return false
}

type IQuery interface {
	IsAnonymous() bool
}

type QueryRoot struct {
	Values []IQuery
}

func (q *QueryRoot) IsAnonymous() bool {
	return false
}

type QueryElem struct {
	anonymous bool
	Key       string
	Values    []IQuery
}

func (q *QueryElem) IsAnonymous() bool {
	return q.anonymous
}

type QueryValue struct {
	Key   string
	Field string
	Value interface{}
}

func (q *QueryValue) IsAnonymous() bool {
	return false
}
