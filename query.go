package gosql

const (
	QueryKey1_and string = "%and"
	QueryKey1_or         = "%or"
	QueryKey1_in         = "%in"
)

const (
	QueryKey2_eq   string = "%eq"   // 等于
	QueryKey2_ne          = "%ne"   // 不等于
	QueryKey2_lt          = "%lt"   // 小于
	QueryKey2_lte         = "%lte"  // 小于等于
	QueryKey2_gt          = "%gt"   // 大于
	QueryKey2_gte         = "%gte"  // 大于等于
	QueryKey2_like        = "%like" // 模糊搜索
)

type IQuery interface {
	QueryKey() string
	QueryValue() []IQuery
	String() string
}

type QueryRoot struct {
	Value IQuery
}

func (q *QueryRoot) ParseJSON(str string) error {
	return nil
}

type QueryElem struct {
	Key   string
	Value IQuery
}

type QueryValue struct {
	Key   string
	Field string
	Value string
}
