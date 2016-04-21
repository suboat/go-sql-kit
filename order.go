package gosql

const (
	OrderKeyASC  string = "+" // 正序
	OrderKeyDESC        = "-" // 反序
)

type IOrder interface {
	IsASC() bool
	IsDESC() bool
}

type OrderRoot struct {
	Value []IOrder
}

func (o *OrderRoot) IsASC() bool {
	return false
}

func (o *OrderRoot) IsDESC() bool {
	return false
}

type OrderValue struct {
	ASC   bool
	Field string
}

func (o *OrderValue) IsASC() bool {
	return o.ASC
}

func (o *OrderValue) IsDESC() bool {
	return !o.ASC
}
